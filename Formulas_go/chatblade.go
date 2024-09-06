package main

// Chatblade - CLI Swiss Army Knife for ChatGPT
// Homepage: https://github.com/npiv/chatblade

import (
	"fmt"
	
	"os/exec"
)

func installChatblade() {
	// Método 1: Descargar y extraer .tar.gz
	chatblade_tar_url := "https://files.pythonhosted.org/packages/7b/69/08b816a91867bc9895e7792cee6a5aa7dd826896340141093b0afdeb2c3d/chatblade-0.6.4.tar.gz"
	chatblade_cmd_tar := exec.Command("curl", "-L", chatblade_tar_url, "-o", "package.tar.gz")
	err := chatblade_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	chatblade_zip_url := "https://files.pythonhosted.org/packages/7b/69/08b816a91867bc9895e7792cee6a5aa7dd826896340141093b0afdeb2c3d/chatblade-0.6.4.zip"
	chatblade_cmd_zip := exec.Command("curl", "-L", chatblade_zip_url, "-o", "package.zip")
	err = chatblade_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	chatblade_bin_url := "https://files.pythonhosted.org/packages/7b/69/08b816a91867bc9895e7792cee6a5aa7dd826896340141093b0afdeb2c3d/chatblade-0.6.4.bin"
	chatblade_cmd_bin := exec.Command("curl", "-L", chatblade_bin_url, "-o", "binary.bin")
	err = chatblade_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	chatblade_src_url := "https://files.pythonhosted.org/packages/7b/69/08b816a91867bc9895e7792cee6a5aa7dd826896340141093b0afdeb2c3d/chatblade-0.6.4.src.tar.gz"
	chatblade_cmd_src := exec.Command("curl", "-L", chatblade_src_url, "-o", "source.tar.gz")
	err = chatblade_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	chatblade_cmd_direct := exec.Command("./binary")
	err = chatblade_cmd_direct.Run()
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
