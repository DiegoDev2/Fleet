package main

// Mallet - MAchine Learning for LanguagE Toolkit
// Homepage: https://mimno.github.io/Mallet/index

import (
	"fmt"
	
	"os/exec"
)

func installMallet() {
	// Método 1: Descargar y extraer .tar.gz
	mallet_tar_url := "https://github.com/mimno/Mallet/releases/download/v202108/Mallet-202108-bin.zip"
	mallet_cmd_tar := exec.Command("curl", "-L", mallet_tar_url, "-o", "package.tar.gz")
	err := mallet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mallet_zip_url := "https://github.com/mimno/Mallet/releases/download/v202108/Mallet-202108-bin.zip"
	mallet_cmd_zip := exec.Command("curl", "-L", mallet_zip_url, "-o", "package.zip")
	err = mallet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mallet_bin_url := "https://github.com/mimno/Mallet/releases/download/v202108/Mallet-202108-bin.zip"
	mallet_cmd_bin := exec.Command("curl", "-L", mallet_bin_url, "-o", "binary.bin")
	err = mallet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mallet_src_url := "https://github.com/mimno/Mallet/releases/download/v202108/Mallet-202108-bin.zip"
	mallet_cmd_src := exec.Command("curl", "-L", mallet_src_url, "-o", "source.tar.gz")
	err = mallet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mallet_cmd_direct := exec.Command("./binary")
	err = mallet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
