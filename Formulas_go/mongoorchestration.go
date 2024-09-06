package main

// MongoOrchestration - REST API to manage MongoDB configurations on a single host
// Homepage: https://github.com/10gen/mongo-orchestration

import (
	"fmt"
	
	"os/exec"
)

func installMongoOrchestration() {
	// Método 1: Descargar y extraer .tar.gz
	mongoorchestration_tar_url := "https://files.pythonhosted.org/packages/f1/c6/3ccc6baa1693168052dff0a96450ac64ea738249c96d890c12c48e4b76a6/mongo_orchestration-0.9.0.tar.gz"
	mongoorchestration_cmd_tar := exec.Command("curl", "-L", mongoorchestration_tar_url, "-o", "package.tar.gz")
	err := mongoorchestration_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mongoorchestration_zip_url := "https://files.pythonhosted.org/packages/f1/c6/3ccc6baa1693168052dff0a96450ac64ea738249c96d890c12c48e4b76a6/mongo_orchestration-0.9.0.zip"
	mongoorchestration_cmd_zip := exec.Command("curl", "-L", mongoorchestration_zip_url, "-o", "package.zip")
	err = mongoorchestration_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mongoorchestration_bin_url := "https://files.pythonhosted.org/packages/f1/c6/3ccc6baa1693168052dff0a96450ac64ea738249c96d890c12c48e4b76a6/mongo_orchestration-0.9.0.bin"
	mongoorchestration_cmd_bin := exec.Command("curl", "-L", mongoorchestration_bin_url, "-o", "binary.bin")
	err = mongoorchestration_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mongoorchestration_src_url := "https://files.pythonhosted.org/packages/f1/c6/3ccc6baa1693168052dff0a96450ac64ea738249c96d890c12c48e4b76a6/mongo_orchestration-0.9.0.src.tar.gz"
	mongoorchestration_cmd_src := exec.Command("curl", "-L", mongoorchestration_src_url, "-o", "source.tar.gz")
	err = mongoorchestration_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mongoorchestration_cmd_direct := exec.Command("./binary")
	err = mongoorchestration_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
