package main

// TRec - Blazingly fast terminal recorder that generates animated gif images for the web
// Homepage: https://github.com/sassman/t-rec-rs

import (
	"fmt"
	
	"os/exec"
)

func installTRec() {
	// Método 1: Descargar y extraer .tar.gz
	trec_tar_url := "https://github.com/sassman/t-rec-rs/archive/refs/tags/v0.7.6.tar.gz"
	trec_cmd_tar := exec.Command("curl", "-L", trec_tar_url, "-o", "package.tar.gz")
	err := trec_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	trec_zip_url := "https://github.com/sassman/t-rec-rs/archive/refs/tags/v0.7.6.zip"
	trec_cmd_zip := exec.Command("curl", "-L", trec_zip_url, "-o", "package.zip")
	err = trec_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	trec_bin_url := "https://github.com/sassman/t-rec-rs/archive/refs/tags/v0.7.6.bin"
	trec_cmd_bin := exec.Command("curl", "-L", trec_bin_url, "-o", "binary.bin")
	err = trec_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	trec_src_url := "https://github.com/sassman/t-rec-rs/archive/refs/tags/v0.7.6.src.tar.gz"
	trec_cmd_src := exec.Command("curl", "-L", trec_src_url, "-o", "source.tar.gz")
	err = trec_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	trec_cmd_direct := exec.Command("./binary")
	err = trec_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: imagemagick")
exec.Command("latte", "install", "imagemagick")
}
