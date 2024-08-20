package main

import (
	//"net/http"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"

	
	"net/http"
	"strings"
	"io"
	"fmt"
	"log"
	"encoding/json"
	"errors"
)

type ProductHandler struct{
	Factory *ProductFactory
}

func NewProductHandler(db *mongo.Client) *ProductHandler {
	return &ProductHandler {
		Factory: NewProductFactory(),
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Message)
}

//Registrar Produto
func InsertProduct(p *Product) error {
	
	collection, ctx, cancel, client, err := ConnectMongoDB()
	
	defer cancel()
    
  	defer func() {
    		if err := client.Disconnect(ctx); err != nil {
    		        panic(err)
     	   	}
	}()
	
	filter := bson.M{"name": p.Name}
	
	var produtoExistente Product
	
	err = collection.FindOne(ctx, filter).Decode(&produtoExistente)
	if err != mongo.ErrNoDocuments {
		if err == nil {
			return errors.New("Product with same name already in database")
		}
		return err
	}

    	_, err = collection.InsertOne(ctx, p)
   	 if err != nil {
        	return err
    	}
	
    	return nil
}

func InsertProductHandler(w http.ResponseWriter, r *http.Request) {
	handler := &ProductHandler{Factory: &ProductFactory{}}
	handler.Register(w, r)
}


func (handler *ProductHandler) Register(w http.ResponseWriter, r *http.Request) {

	body, _ := io.ReadAll(r.Body)

	product, err := handler.Factory.CreateFromJSON(body)
	
	if err != nil {
		http.Error(w, "Invalid json body: " + err.Error(), http.StatusBadRequest)
		return
	}

	err = InsertProduct(product) 

	if err != nil {
		http.Error(w, "Unable to insert product in the database", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product inserted successfully"})
}

//Listar Produtos
func ListProducts() ([]Product, error) {
	collection, ctx, cancel, client, err := ConnectMongoDB()
	
	defer cancel()
	
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect to database: %v", err)
		}
	}()
	
	cursor, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Printf("Falha ao achar documentos em 'products'")
		return nil, fmt.Errorf("Failed to find documents in collection 'products': %v", err)
	}
	defer cursor.Close(ctx)
	
	var products []Product
	
	for cursor.Next(ctx) {
		var product Product
		if err := cursor.Decode(&product); err != nil {
			log.Printf("Falha ao decodificar produto")
			return nil, fmt.Errorf("Failed to decode product: %v", err)
		}
		
		products = append(products, product)
	}
	
	if err := cursor.Err(); err != nil {
		log.Printf("Falha no cursor")
		return nil, fmt.Errorf("Cursor error: %v", err)
	}
	log.Printf("Showing products.")
	return products, nil
}

func ListProductsHandler(w http.ResponseWriter, r *http.Request) {
	
	products, err := ListProducts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

//Deletar produtos
func DeleteAllProductByName(name string) error {
	collection, ctx, cancel, client, err := ConnectMongoDB()

	defer cancel()
		
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Fatalf("Failed to disconnect to database: %v", err)
		}
	}()
	
	var product Product
	filter := bson.D{{Key: "name", Value: name}}
	
	//verifica se produto está disponivel
	err = collection.FindOne(ctx, filter).Decode(&product)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("No document found with name: %s", name)
			return fmt.Errorf("No product found with name: %s", name)
		}
		
		log.Printf("Failed to retrieve document from 'products': %v", err)
		return fmt.Errorf("Failed to retrieve product: %v", err)
	}
	
	if !product.Disponibility {
		log.Printf("Product with name: '%s' is not available:", name)
		return fmt.Errorf("Product with name: '%s' is not available", name)
	}
	
	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		log.Printf("Failed to delete document from 'products'")
		return fmt.Errorf("Failed to delete product: %v", err)
	}
	
	if result.DeletedCount == 0 {
		log.Printf("No document found with name: %s", name)
		return fmt.Errorf("No product found with name: %s", name)
	}
	
	log.Printf("Deleted product with name: %s", name)
	return nil
}

func DeleteProductHandler(w http.ResponseWriter, r *http.Request) {

	name := r.URL.Query().Get("name")
	if name == "" {
		http.Error(w, "Missing product name in query", http.StatusBadRequest)
		return
	}

	err := DeleteAllProductByName(name)
	if err != nil {
		http.Error(w, "Unable to delete product: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Product deleted successfully"})
}


func UpdateProductHandler(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "Missing product id in query", http.StatusBadRequest)
        return
    }
    log.Printf("Received product id: %s", id) // Log para verificar o id recebido

    err := UpdateProduct(r, id)
    if err != nil {
        http.Error(w, "Unable to update product: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]string{"message": "Product updated successfully"})
}

func UpdateProduct(r *http.Request, id string) error {
    var updateData map[string]interface{}
    decoder := json.NewDecoder(r.Body)

    err := decoder.Decode(&updateData)
    if err != nil {
        log.Printf("Unable to decode body")
        return fmt.Errorf("Invalid request payload: %v", err)
    }
    defer r.Body.Close()

    // Verificar se há um ID válido
    if id == "" {
        log.Printf("Empty product ID")
        return fmt.Errorf("Product ID is required")
    }
	
    if weight, ok := updateData["weight"].(string); ok {
    	if !strings.HasSuffix(weight, "g") { 
        	updateData["weight"] = weight + "g"
    	}
    }

    // Lógica para atualizar a disponibilidade baseada na quantidade
    if amount, ok := updateData["amount"].(float64); ok {
        if amount == 0 {
            updateData["disponibility"] = false
        } else {
            updateData["disponibility"] = true
        }
    }

    // Conectar ao MongoDB e preparar para atualizar o produto
    collection, ctx, cancel, client, err := ConnectMongoDB()
    defer cancel()
    defer func() {
        if err := client.Disconnect(ctx); err != nil {
            log.Fatalf("Failed to disconnect from the database: %v", err)
        }
    }()

    // Filtrar pelo ID do produto
    filter := bson.D{{Key: "id", Value: id}}

    // Construir a atualização
    update := bson.D{{"$set", updateData}}

    // Executar a atualização no MongoDB
    result, err := collection.UpdateOne(ctx, filter, update)
    if err != nil {
        log.Printf("Failed to update product: %v", err)
        return fmt.Errorf("Failed to update product: %v", err)
    }

    // Verificar se o produto foi encontrado e atualizado
    if result.MatchedCount == 0 {
        log.Printf("No product found with ID: %s", id)
        return fmt.Errorf("No product found with ID: %s", id)
    }

    log.Printf("Updated product with ID: %s", id)
    return nil
}


func GetProductByName(productName string) (*Product, error) {
    collection, ctx, cancel, client, err := ConnectMongoDB()
    if err != nil {
        return nil, fmt.Errorf("Failed to connect to MongoDB: %v", err)
    }
    defer cancel()
    defer func() {
        if err := client.Disconnect(ctx); err != nil {
            log.Fatalf("Failed to disconnect from the database: %v", err)
        }
    }()

    var product Product
    err = collection.FindOne(ctx, bson.M{"name": productName}).Decode(&product)
    if err != nil {
        if errors.Is(err, mongo.ErrNoDocuments) {
            return nil, fmt.Errorf("Product not found")
        }
        return nil, fmt.Errorf("Failed to find product: %v", err)
    }

    return &product, nil
}

func GetProductByNameHandler(w http.ResponseWriter, r *http.Request) {
    productName := r.URL.Query().Get("name") // Supondo que você passe o nome como parâmetro de query "name"
    if productName == "" {
        http.Error(w, "Product name parameter is required", http.StatusBadRequest)
        return
    }

    product, err := GetProductByName(productName)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to get product: %v", err), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(product)
}








