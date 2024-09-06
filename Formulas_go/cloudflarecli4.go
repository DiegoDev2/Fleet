package main

// CloudflareCli4 - CLI for Cloudflare API v4
// Homepage: https://github.com/cloudflare/python-cloudflare-cli4

import (
	"fmt"
	
	"os/exec"
)

func installCloudflareCli4() {
	// Método 1: Descargar y extraer .tar.gz
	cloudflarecli4_tar_url := "https://github.com/cloudflare/python-cloudflare-cli4/archive/refs/tags/2.19.4.tar.gz"
	cloudflarecli4_cmd_tar := exec.Command("curl", "-L", cloudflarecli4_tar_url, "-o", "package.tar.gz")
	err := cloudflarecli4_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cloudflarecli4_zip_url := "https://github.com/cloudflare/python-cloudflare-cli4/archive/refs/tags/2.19.4.zip"
	cloudflarecli4_cmd_zip := exec.Command("curl", "-L", cloudflarecli4_zip_url, "-o", "package.zip")
	err = cloudflarecli4_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cloudflarecli4_bin_url := "https://github.com/cloudflare/python-cloudflare-cli4/archive/refs/tags/2.19.4.bin"
	cloudflarecli4_cmd_bin := exec.Command("curl", "-L", cloudflarecli4_bin_url, "-o", "binary.bin")
	err = cloudflarecli4_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cloudflarecli4_src_url := "https://github.com/cloudflare/python-cloudflare-cli4/archive/refs/tags/2.19.4.src.tar.gz"
	cloudflarecli4_cmd_src := exec.Command("curl", "-L", cloudflarecli4_src_url, "-o", "source.tar.gz")
	err = cloudflarecli4_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cloudflarecli4_cmd_direct := exec.Command("./binary")
	err = cloudflarecli4_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
