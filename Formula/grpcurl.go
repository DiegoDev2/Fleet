package main

// Grpcurl - Like cURL, but for gRPC
// Homepage: https://github.com/fullstorydev/grpcurl

import (
	"fmt"
	
	"os/exec"
)

func installGrpcurl() {
	// Método 1: Descargar y extraer .tar.gz
	grpcurl_tar_url := "https://github.com/fullstorydev/grpcurl/archive/refs/tags/v1.9.1.tar.gz"
	grpcurl_cmd_tar := exec.Command("curl", "-L", grpcurl_tar_url, "-o", "package.tar.gz")
	err := grpcurl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grpcurl_zip_url := "https://github.com/fullstorydev/grpcurl/archive/refs/tags/v1.9.1.zip"
	grpcurl_cmd_zip := exec.Command("curl", "-L", grpcurl_zip_url, "-o", "package.zip")
	err = grpcurl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grpcurl_bin_url := "https://github.com/fullstorydev/grpcurl/archive/refs/tags/v1.9.1.bin"
	grpcurl_cmd_bin := exec.Command("curl", "-L", grpcurl_bin_url, "-o", "binary.bin")
	err = grpcurl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grpcurl_src_url := "https://github.com/fullstorydev/grpcurl/archive/refs/tags/v1.9.1.src.tar.gz"
	grpcurl_cmd_src := exec.Command("curl", "-L", grpcurl_src_url, "-o", "source.tar.gz")
	err = grpcurl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grpcurl_cmd_direct := exec.Command("./binary")
	err = grpcurl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
