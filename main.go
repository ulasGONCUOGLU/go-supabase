package main

import (
	"fmt"
	"strings"

	"github.com/supabase-community/supabase-go"
)

func main() {
	supabaseUrl := "https://vgzkgfiwxxzjicaenpof.supabase.co"
	supabaseKey := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6InZnemtnZml3eHh6amljYWVucG9mIiwicm9sZSI6ImFub24iLCJpYXQiOjE3Mzk0NTE4NjYsImV4cCI6MjA1NTAyNzg2Nn0.r992f__ra07NHIuReYy_Hdv3GPpqJwRSM2gcauGlFf8"

	
	ulas, err := supabase.NewClient(supabaseUrl, supabaseKey, &supabase.ClientOptions{})
	if err != nil {
		fmt.Println("cannot initalize client", err)
	}
	
	
	/*
	data, count, err := ulas.From("logs").Select("*", "exact", false).Execute()

	fmt.Println("Veri:", data)
	
	fmt.Println("Adet:", count)
	*/


		
	rooms, _, err := ulas.From("logs").Select("*", "", false).ExecuteString()
	if err != nil {
		panic(err)
	}

	// fmt.Println(strings.Contains(rooms, "id"))
	// fmt.Println(strings.Trim(rooms, "[]{},"))
	fmt.Println(strings.Split(rooms, ","))
}
