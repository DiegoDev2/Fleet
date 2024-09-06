package main

// Uvicorn - ASGI web server
// Homepage: https://www.uvicorn.org/

import (
	"fmt"
	
	"os/exec"
)

func installUvicorn() {
	// Método 1: Descargar y extraer .tar.gz
	uvicorn_tar_url := "https://files.pythonhosted.org/packages/5a/01/5e637e7aa9dd031be5376b9fb749ec20b86f5a5b6a49b87fabd374d5fa9f/uvicorn-0.30.6.tar.gz"
	uvicorn_cmd_tar := exec.Command("curl", "-L", uvicorn_tar_url, "-o", "package.tar.gz")
	err := uvicorn_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uvicorn_zip_url := "https://files.pythonhosted.org/packages/5a/01/5e637e7aa9dd031be5376b9fb749ec20b86f5a5b6a49b87fabd374d5fa9f/uvicorn-0.30.6.zip"
	uvicorn_cmd_zip := exec.Command("curl", "-L", uvicorn_zip_url, "-o", "package.zip")
	err = uvicorn_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uvicorn_bin_url := "https://files.pythonhosted.org/packages/5a/01/5e637e7aa9dd031be5376b9fb749ec20b86f5a5b6a49b87fabd374d5fa9f/uvicorn-0.30.6.bin"
	uvicorn_cmd_bin := exec.Command("curl", "-L", uvicorn_bin_url, "-o", "binary.bin")
	err = uvicorn_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uvicorn_src_url := "https://files.pythonhosted.org/packages/5a/01/5e637e7aa9dd031be5376b9fb749ec20b86f5a5b6a49b87fabd374d5fa9f/uvicorn-0.30.6.src.tar.gz"
	uvicorn_cmd_src := exec.Command("curl", "-L", uvicorn_src_url, "-o", "source.tar.gz")
	err = uvicorn_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uvicorn_cmd_direct := exec.Command("./binary")
	err = uvicorn_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
