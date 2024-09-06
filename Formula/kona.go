package main

// Kona - Open-source implementation of the K programming language
// Homepage: https://github.com/kevinlawler/kona

import (
	"fmt"
	
	"os/exec"
)

func installKona() {
	// Método 1: Descargar y extraer .tar.gz
	kona_tar_url := "https://github.com/kevinlawler/kona/archive/refs/tags/Win64-20211225.tar.gz"
	kona_cmd_tar := exec.Command("curl", "-L", kona_tar_url, "-o", "package.tar.gz")
	err := kona_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kona_zip_url := "https://github.com/kevinlawler/kona/archive/refs/tags/Win64-20211225.zip"
	kona_cmd_zip := exec.Command("curl", "-L", kona_zip_url, "-o", "package.zip")
	err = kona_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kona_bin_url := "https://github.com/kevinlawler/kona/archive/refs/tags/Win64-20211225.bin"
	kona_cmd_bin := exec.Command("curl", "-L", kona_bin_url, "-o", "binary.bin")
	err = kona_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kona_src_url := "https://github.com/kevinlawler/kona/archive/refs/tags/Win64-20211225.src.tar.gz"
	kona_cmd_src := exec.Command("curl", "-L", kona_src_url, "-o", "source.tar.gz")
	err = kona_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kona_cmd_direct := exec.Command("./binary")
	err = kona_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
