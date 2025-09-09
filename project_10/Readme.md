* Question : Write a program to Build a configuration loader struct from JSON/YAML.

==> lets understand what is JSON file
    : JSON is a javaScript object notation file
    : Its a light weight text format for storing and exchanging data 
    : Uses key-value pairs

    for example :- {
                        "app_name": "MyApp",
                        "debug": true,
                        "server": {
                        "host": "127.0.0.1",
                        "port": 8080
                            }
                    }

==> What is YAML file
    : YAML stand for Yet Another MarkUp Language
    : Instead, it is designed as a data serialization language → a way to represent structured data (like configs) in a human-readable format.

    for example :- app_name: "MyApp"
                   debug: true
                   server:
                    host: "127.0.0.1"
                    port: 8080


* How do we use struct 
==> Here we use two type of struct that is 
                                            1.) Named struct == normal struct
                                            2.) Inline struct  == type of nested struct without a seprate name
        

Various ways to declare a inline struct

                    1.) Inline Struct inside a variable
                    -> Here we declare a struct inside a variable 
                    that is :
                                person := struct {
                                                    Name string
                                                    Age  int
                                                }{
                                                    Name: "Alice",
                                                    Age:  25,
                                                }
 
                    2.) Inline Struct inside another Struct
                    -> Instead of creating a new type separately, you directly define a small struct as a field inside another struct.
                    that is :
                                type Config struct {
                                                    AppName string
                                                     Debug   bool

                                                    // Inline struct as a field
                                                    Server struct {
                                                        Host string
                                                        Port int
                                                        }
                                                    }

* When to use Inline struct
When to Use
            1.)One-time use: The nested struct is only needed inside this parent struct.
            2.)Small data: The nested struct has only a few fields.
            3.)Tight coupling: The nested struct has no meaning outside the parent.
            4.)Config/data models: Common in JSON/YAML configs, where you just mirror the file structure.



                            *   Inline struct = Quick, local, disposable.
                            *   Named struct = Clean, reusable, modular.

* Why we use inline struct in loader file code

==> Feature	                            Inline Struct	                          Named Struct

Code length	                        Shorter, compact	                         More verbose

Readability	                        Mirrors JSON/YAML directly	                 Requires jumping between types

Reusability	                        ❌ Cannot reuse	                            ✅ Can reuse across program

Testing	                            ❌ Hard to test separately	                Can test sub-struct independently

Good for	                        Small configs, one-time use	                Large configs,reusable components


==> We know json/yaml file are used once where it is neccesary to make it is reuseable

we have take a config.json file as an example