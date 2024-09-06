package main

// AwsEsProxy - Small proxy between HTTP client and AWS Elasticsearch
// Homepage: https://github.com/abutaha/aws-es-proxy

import (
	"fmt"
	
	"os/exec"
)

func installAwsEsProxy() {
	// Método 1: Descargar y extraer .tar.gz
	awsesproxy_tar_url := "https://github.com/abutaha/aws-es-proxy/archive/refs/tags/v1.5.tar.gz"
	awsesproxy_cmd_tar := exec.Command("curl", "-L", awsesproxy_tar_url, "-o", "package.tar.gz")
	err := awsesproxy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	awsesproxy_zip_url := "https://github.com/abutaha/aws-es-proxy/archive/refs/tags/v1.5.zip"
	awsesproxy_cmd_zip := exec.Command("curl", "-L", awsesproxy_zip_url, "-o", "package.zip")
	err = awsesproxy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	awsesproxy_bin_url := "https://github.com/abutaha/aws-es-proxy/archive/refs/tags/v1.5.bin"
	awsesproxy_cmd_bin := exec.Command("curl", "-L", awsesproxy_bin_url, "-o", "binary.bin")
	err = awsesproxy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	awsesproxy_src_url := "https://github.com/abutaha/aws-es-proxy/archive/refs/tags/v1.5.src.tar.gz"
	awsesproxy_cmd_src := exec.Command("curl", "-L", awsesproxy_src_url, "-o", "source.tar.gz")
	err = awsesproxy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	awsesproxy_cmd_direct := exec.Command("./binary")
	err = awsesproxy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
