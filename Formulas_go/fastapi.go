package main

// Fastapi - CLI for FastAPI framework
// Homepage: https://fastapi.tiangolo.com/

import (
	"fmt"
	
	"os/exec"
)

func installFastapi() {
	// Método 1: Descargar y extraer .tar.gz
	fastapi_tar_url := "https://files.pythonhosted.org/packages/92/36/2eafe7d5ef26c1ae8e396cfcde3c9e9c50385f5c486dccdfba420d210c3d/fastapi-0.112.4.tar.gz"
	fastapi_cmd_tar := exec.Command("curl", "-L", fastapi_tar_url, "-o", "package.tar.gz")
	err := fastapi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fastapi_zip_url := "https://files.pythonhosted.org/packages/92/36/2eafe7d5ef26c1ae8e396cfcde3c9e9c50385f5c486dccdfba420d210c3d/fastapi-0.112.4.zip"
	fastapi_cmd_zip := exec.Command("curl", "-L", fastapi_zip_url, "-o", "package.zip")
	err = fastapi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fastapi_bin_url := "https://files.pythonhosted.org/packages/92/36/2eafe7d5ef26c1ae8e396cfcde3c9e9c50385f5c486dccdfba420d210c3d/fastapi-0.112.4.bin"
	fastapi_cmd_bin := exec.Command("curl", "-L", fastapi_bin_url, "-o", "binary.bin")
	err = fastapi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fastapi_src_url := "https://files.pythonhosted.org/packages/92/36/2eafe7d5ef26c1ae8e396cfcde3c9e9c50385f5c486dccdfba420d210c3d/fastapi-0.112.4.src.tar.gz"
	fastapi_cmd_src := exec.Command("curl", "-L", fastapi_src_url, "-o", "source.tar.gz")
	err = fastapi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fastapi_cmd_direct := exec.Command("./binary")
	err = fastapi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
