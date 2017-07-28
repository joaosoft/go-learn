package swagger

//This file is generated automatically. Do not try to edit it manually.

var ResourceListingJson = `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "{{.}}",
    "apis": [
        {
            "path": "/products",
            "description": "Product Search"
        },
        {
            "path": "/offers",
            "description": "Product Offers Search"
        }
    ],
    "info": {
        "title": "Sonae Seach and Navigation searvice",
        "description": "Search service",
        "contact": "teste",
        "termsOfServiceUrl": "teste",
        "license": "BSD",
        "licenseUrl": "http://opensource.org/licenses/BSD-2-Clause"
    }
}`
var ApiDescriptionsJson = map[string]string{"products": `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "{{.}}",
    "resourcePath": "/products",
    "produces": [
        "application/json"
    ],
    "apis": [
        {
            "path": "api/products",
            "description": "Get all product in query",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "Get Products",
                    "type": "search-and-navigation.services.search.domain.ProductOfferCollection",
                    "items": {},
                    "summary": "Get all product in query",
                    "parameters": [
                        {
                            "paramType": "query",
                            "name": "q",
                            "description": "Query",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Success",
                            "responseType": "object",
                            "responseModel": "search-and-navigation.services.search.domain.ProductOfferCollection"
                        },
                        {
                            "code": 500,
                            "message": "Unexpected error searching products",
                            "responseType": "object",
                            "responseModel": "search-and-navigation.services.search.vo.ApiError"
                        }
                    ],
                    "produces": [
                        "application/json"
                    ]
                }
            ]
        }
    ],
    "models": {
        "github.com.satori.go.uuid.UUID": {
            "id": "github.com.satori.go.uuid.UUID",
            "properties": null
        },
        "search-and-navigation.common.api.ResponseCollectionPagination": {
            "id": "search-and-navigation.common.api.ResponseCollectionPagination",
            "properties": {
                "page": {
                    "type": "int",
                    "description": "Current page",
                    "items": {},
                    "format": ""
                },
                "per_page": {
                    "type": "int",
                    "description": "Number of items per page",
                    "items": {},
                    "format": ""
                },
                "total_items": {
                    "type": "int",
                    "description": "Total items",
                    "items": {},
                    "format": ""
                },
                "total_pages": {
                    "type": "int",
                    "description": "Number of pages",
                    "items": {},
                    "format": ""
                }
            }
        },
        "search-and-navigation.services.search.domain.ProductOffer": {
            "id": "search-and-navigation.services.search.domain.ProductOffer",
            "properties": {
                "final_price": {
                    "type": "float32",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "is_in_stock": {
                    "type": "bool",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "offer_id": {
                    "type": "github.com.satori.go.uuid.UUID",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "price": {
                    "type": "float32",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "product_id": {
                    "type": "github.com.satori.go.uuid.UUID",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "seller_id": {
                    "type": "github.com.satori.go.uuid.UUID",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "source_id": {
                    "type": "github.com.satori.go.uuid.UUID",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "source_offer_code": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "source_product_code": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "state": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updated_at": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "search-and-navigation.services.search.domain.ProductOfferCollection": {
            "id": "search-and-navigation.services.search.domain.ProductOfferCollection",
            "properties": {
                "offers": {
                    "type": "array",
                    "description": "Offers",
                    "items": {
                        "$ref": "search-and-navigation.services.search.domain.ProductOffer"
                    },
                    "format": ""
                },
                "pagination": {
                    "type": "search-and-navigation.common.api.ResponseCollectionPagination",
                    "description": "Pagination",
                    "items": {},
                    "format": ""
                }
            }
        },
        "search-and-navigation.services.search.vo.ApiError": {
            "id": "search-and-navigation.services.search.vo.ApiError",
            "properties": {
                "code": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "message": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`, "offers": `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "{{.}}",
    "resourcePath": "/offers",
    "produces": [
        "application/json"
    ],
    "apis": [
        {
            "path": "api/products/{productUUID}/channels/{channelUUID}/offers",
            "description": "Get all offers by Product UUID and accepts orders and filters",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "Get Offers By Product UUID",
                    "type": "search-and-navigation.services.search.domain.ProductOfferCollection",
                    "items": {},
                    "summary": "Get all offers by Product UUID and accepts orders and filters",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "productUUID",
                            "description": "Product UUID",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "channelUUID",
                            "description": "Channel UUID",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "offerRequest",
                            "description": "Search parameters",
                            "dataType": "search-and-navigation.services.search.controllers.GetProductOffersByProductUUIDRequest",
                            "type": "search-and-navigation.services.search.controllers.GetProductOffersByProductUUIDRequest",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Success",
                            "responseType": "object",
                            "responseModel": "search-and-navigation.services.search.domain.ProductOfferCollection"
                        },
                        {
                            "code": 400,
                            "message": "Invalid Request Body",
                            "responseType": "object",
                            "responseModel": "search-and-navigation.services.search.vo.ApiError"
                        },
                        {
                            "code": 500,
                            "message": "Unexpected error searching offers",
                            "responseType": "object",
                            "responseModel": "search-and-navigation.services.search.vo.ApiError"
                        }
                    ],
                    "produces": [
                        "application/json"
                    ]
                }
            ]
        },
        {
            "path": "api/sellers/{sellerUUID}/channels/{channelUUID}/offers",
            "description": "Get all offers by Seller UUID and accepts orders and filters",
            "operations": [
                {
                    "httpMethod": "POST",
                    "nickname": "Get Offers By Seller UUID",
                    "type": "search-and-navigation.services.search.domain.ProductOfferCollection",
                    "items": {},
                    "summary": "Get all offers by Seller UUID and accepts orders and filters",
                    "parameters": [
                        {
                            "paramType": "path",
                            "name": "sellerUUID",
                            "description": "Seller UUID",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "path",
                            "name": "channelUUID",
                            "description": "Channel UUID",
                            "dataType": "string",
                            "type": "string",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        },
                        {
                            "paramType": "body",
                            "name": "offerRequest",
                            "description": "Search parameters",
                            "dataType": "search-and-navigation.services.search.controllers.GetProductOffersBySellerUUIDRequest",
                            "type": "search-and-navigation.services.search.controllers.GetProductOffersBySellerUUIDRequest",
                            "format": "",
                            "allowMultiple": false,
                            "required": true,
                            "minimum": 0,
                            "maximum": 0
                        }
                    ],
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "Success",
                            "responseType": "object",
                            "responseModel": "search-and-navigation.services.search.domain.ProductOfferCollection"
                        },
                        {
                            "code": 400,
                            "message": "Invalid Request Body",
                            "responseType": "object",
                            "responseModel": "search-and-navigation.services.search.vo.ApiError"
                        },
                        {
                            "code": 500,
                            "message": "Unexpected error searching offers",
                            "responseType": "object",
                            "responseModel": "search-and-navigation.services.search.vo.ApiError"
                        }
                    ],
                    "produces": [
                        "application/json"
                    ]
                }
            ]
        }
    ],
    "models": {
        "github.com.satori.go.uuid.UUID": {
            "id": "github.com.satori.go.uuid.UUID",
            "properties": null
        },
        "search-and-navigation.common.api.Pagination": {
            "id": "search-and-navigation.common.api.Pagination",
            "properties": {
                "page": {
                    "type": "int",
                    "description": "Current page",
                    "items": {},
                    "format": ""
                },
                "per_page": {
                    "type": "int",
                    "description": "Number of items per page",
                    "items": {},
                    "format": ""
                }
            }
        },
        "search-and-navigation.common.api.ResponseCollectionPagination": {
            "id": "search-and-navigation.common.api.ResponseCollectionPagination",
            "properties": {
                "page": {
                    "type": "int",
                    "description": "Current page",
                    "items": {},
                    "format": ""
                },
                "per_page": {
                    "type": "int",
                    "description": "Number of items per page",
                    "items": {},
                    "format": ""
                },
                "total_items": {
                    "type": "int",
                    "description": "Total items",
                    "items": {},
                    "format": ""
                },
                "total_pages": {
                    "type": "int",
                    "description": "Number of pages",
                    "items": {},
                    "format": ""
                }
            }
        },
        "search-and-navigation.services.search.controllers.GetProductOffersByProductUUIDRequest": {
            "id": "search-and-navigation.services.search.controllers.GetProductOffersByProductUUIDRequest",
            "properties": {
                "order": {
                    "type": "array",
                    "description": "Sorting parameters",
                    "items": {
                        "type": "string"
                    },
                    "format": ""
                },
                "pagination": {
                    "type": "search-and-navigation.common.api.Pagination",
                    "description": "Pagination",
                    "items": {},
                    "format": ""
                }
            }
        },
        "search-and-navigation.services.search.controllers.GetProductOffersBySellerUUIDRequest": {
            "id": "search-and-navigation.services.search.controllers.GetProductOffersBySellerUUIDRequest",
            "properties": {
                "order": {
                    "type": "array",
                    "description": "Sorting parameters",
                    "items": {
                        "type": "string"
                    },
                    "format": ""
                },
                "pagination": {
                    "type": "search-and-navigation.common.api.Pagination",
                    "description": "Pagination",
                    "items": {},
                    "format": ""
                }
            }
        },
        "search-and-navigation.services.search.domain.ProductOffer": {
            "id": "search-and-navigation.services.search.domain.ProductOffer",
            "properties": {
                "final_price": {
                    "type": "float32",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "is_in_stock": {
                    "type": "bool",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "offer_id": {
                    "type": "github.com.satori.go.uuid.UUID",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "price": {
                    "type": "float32",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "product_id": {
                    "type": "github.com.satori.go.uuid.UUID",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "seller_id": {
                    "type": "github.com.satori.go.uuid.UUID",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "source_id": {
                    "type": "github.com.satori.go.uuid.UUID",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "source_offer_code": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "source_product_code": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "state": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "updated_at": {
                    "type": "Time",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        },
        "search-and-navigation.services.search.domain.ProductOfferCollection": {
            "id": "search-and-navigation.services.search.domain.ProductOfferCollection",
            "properties": {
                "offers": {
                    "type": "array",
                    "description": "Offers",
                    "items": {
                        "$ref": "search-and-navigation.services.search.domain.ProductOffer"
                    },
                    "format": ""
                },
                "pagination": {
                    "type": "search-and-navigation.common.api.ResponseCollectionPagination",
                    "description": "Pagination",
                    "items": {},
                    "format": ""
                }
            }
        },
        "search-and-navigation.services.search.vo.ApiError": {
            "id": "search-and-navigation.services.search.vo.ApiError",
            "properties": {
                "code": {
                    "type": "int64",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "message": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`}
