package main

// S3scanner - Scan for misconfigured S3 buckets across S3-compatible APIs!
// Homepage: https://github.com/sa7mon/S3Scanner

import (
	"fmt"
	
	"os/exec"
)

func installS3scanner() {
	// Método 1: Descargar y extraer .tar.gz
	s3scanner_tar_url := "https://github.com/sa7mon/S3Scanner/archive/refs/tags/v3.0.4.tar.gz"
	s3scanner_cmd_tar := exec.Command("curl", "-L", s3scanner_tar_url, "-o", "package.tar.gz")
	err := s3scanner_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	s3scanner_zip_url := "https://github.com/sa7mon/S3Scanner/archive/refs/tags/v3.0.4.zip"
	s3scanner_cmd_zip := exec.Command("curl", "-L", s3scanner_zip_url, "-o", "package.zip")
	err = s3scanner_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	s3scanner_bin_url := "https://github.com/sa7mon/S3Scanner/archive/refs/tags/v3.0.4.bin"
	s3scanner_cmd_bin := exec.Command("curl", "-L", s3scanner_bin_url, "-o", "binary.bin")
	err = s3scanner_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	s3scanner_src_url := "https://github.com/sa7mon/S3Scanner/archive/refs/tags/v3.0.4.src.tar.gz"
	s3scanner_cmd_src := exec.Command("curl", "-L", s3scanner_src_url, "-o", "source.tar.gz")
	err = s3scanner_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	s3scanner_cmd_direct := exec.Command("./binary")
	err = s3scanner_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
