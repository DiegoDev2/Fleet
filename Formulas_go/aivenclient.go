package main

// AivenClient - Official command-line client for Aiven
// Homepage: https://docs.aiven.io/docs/tools/cli

import (
	"fmt"
	
	"os/exec"
)

func installAivenClient() {
	// Método 1: Descargar y extraer .tar.gz
	aivenclient_tar_url := "https://files.pythonhosted.org/packages/fd/45/9d082f86318981c254f5c784b0dda9d673cc30f992cc4f6bf69cecb08aec/aiven_client-4.2.1.tar.gz"
	aivenclient_cmd_tar := exec.Command("curl", "-L", aivenclient_tar_url, "-o", "package.tar.gz")
	err := aivenclient_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	aivenclient_zip_url := "https://files.pythonhosted.org/packages/fd/45/9d082f86318981c254f5c784b0dda9d673cc30f992cc4f6bf69cecb08aec/aiven_client-4.2.1.zip"
	aivenclient_cmd_zip := exec.Command("curl", "-L", aivenclient_zip_url, "-o", "package.zip")
	err = aivenclient_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	aivenclient_bin_url := "https://files.pythonhosted.org/packages/fd/45/9d082f86318981c254f5c784b0dda9d673cc30f992cc4f6bf69cecb08aec/aiven_client-4.2.1.bin"
	aivenclient_cmd_bin := exec.Command("curl", "-L", aivenclient_bin_url, "-o", "binary.bin")
	err = aivenclient_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	aivenclient_src_url := "https://files.pythonhosted.org/packages/fd/45/9d082f86318981c254f5c784b0dda9d673cc30f992cc4f6bf69cecb08aec/aiven_client-4.2.1.src.tar.gz"
	aivenclient_cmd_src := exec.Command("curl", "-L", aivenclient_src_url, "-o", "source.tar.gz")
	err = aivenclient_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	aivenclient_cmd_direct := exec.Command("./binary")
	err = aivenclient_cmd_direct.Run()
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
