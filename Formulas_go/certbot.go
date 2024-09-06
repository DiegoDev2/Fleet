package main

// Certbot - Tool to obtain certs from Let's Encrypt and autoenable HTTPS
// Homepage: https://certbot.eff.org/

import (
	"fmt"
	
	"os/exec"
)

func installCertbot() {
	// Método 1: Descargar y extraer .tar.gz
	certbot_tar_url := "https://files.pythonhosted.org/packages/c0/e9/5e637d66a9fe6d93c4c075d539a3b949244a9b795b3f07a0998951af9b00/certbot-2.11.0.tar.gz"
	certbot_cmd_tar := exec.Command("curl", "-L", certbot_tar_url, "-o", "package.tar.gz")
	err := certbot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	certbot_zip_url := "https://files.pythonhosted.org/packages/c0/e9/5e637d66a9fe6d93c4c075d539a3b949244a9b795b3f07a0998951af9b00/certbot-2.11.0.zip"
	certbot_cmd_zip := exec.Command("curl", "-L", certbot_zip_url, "-o", "package.zip")
	err = certbot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	certbot_bin_url := "https://files.pythonhosted.org/packages/c0/e9/5e637d66a9fe6d93c4c075d539a3b949244a9b795b3f07a0998951af9b00/certbot-2.11.0.bin"
	certbot_cmd_bin := exec.Command("curl", "-L", certbot_bin_url, "-o", "binary.bin")
	err = certbot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	certbot_src_url := "https://files.pythonhosted.org/packages/c0/e9/5e637d66a9fe6d93c4c075d539a3b949244a9b795b3f07a0998951af9b00/certbot-2.11.0.src.tar.gz"
	certbot_cmd_src := exec.Command("curl", "-L", certbot_src_url, "-o", "source.tar.gz")
	err = certbot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	certbot_cmd_direct := exec.Command("./binary")
	err = certbot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: augeas")
exec.Command("latte", "install", "augeas")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: dialog")
exec.Command("latte", "install", "dialog")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
