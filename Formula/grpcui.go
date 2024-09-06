package main

// Grpcui - Interactive web UI for gRPC, along the lines of postman
// Homepage: https://github.com/fullstorydev/grpcui

import (
	"fmt"
	
	"os/exec"
)

func installGrpcui() {
	// Método 1: Descargar y extraer .tar.gz
	grpcui_tar_url := "https://github.com/fullstorydev/grpcui/archive/refs/tags/v1.4.1.tar.gz"
	grpcui_cmd_tar := exec.Command("curl", "-L", grpcui_tar_url, "-o", "package.tar.gz")
	err := grpcui_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grpcui_zip_url := "https://github.com/fullstorydev/grpcui/archive/refs/tags/v1.4.1.zip"
	grpcui_cmd_zip := exec.Command("curl", "-L", grpcui_zip_url, "-o", "package.zip")
	err = grpcui_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grpcui_bin_url := "https://github.com/fullstorydev/grpcui/archive/refs/tags/v1.4.1.bin"
	grpcui_cmd_bin := exec.Command("curl", "-L", grpcui_bin_url, "-o", "binary.bin")
	err = grpcui_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grpcui_src_url := "https://github.com/fullstorydev/grpcui/archive/refs/tags/v1.4.1.src.tar.gz"
	grpcui_cmd_src := exec.Command("curl", "-L", grpcui_src_url, "-o", "source.tar.gz")
	err = grpcui_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grpcui_cmd_direct := exec.Command("./binary")
	err = grpcui_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
}
