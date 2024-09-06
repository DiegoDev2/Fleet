package main

// Jrnl - Command-line note taker
// Homepage: https://jrnl.sh/en/stable/

import (
	"fmt"
	
	"os/exec"
)

func installJrnl() {
	// Método 1: Descargar y extraer .tar.gz
	jrnl_tar_url := "https://files.pythonhosted.org/packages/b6/65/3b0649ac261e3cf7c110acbdd74b13eeb3e9e6a91eb41832cb4d7d1f9049/jrnl-4.1.tar.gz"
	jrnl_cmd_tar := exec.Command("curl", "-L", jrnl_tar_url, "-o", "package.tar.gz")
	err := jrnl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jrnl_zip_url := "https://files.pythonhosted.org/packages/b6/65/3b0649ac261e3cf7c110acbdd74b13eeb3e9e6a91eb41832cb4d7d1f9049/jrnl-4.1.zip"
	jrnl_cmd_zip := exec.Command("curl", "-L", jrnl_zip_url, "-o", "package.zip")
	err = jrnl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jrnl_bin_url := "https://files.pythonhosted.org/packages/b6/65/3b0649ac261e3cf7c110acbdd74b13eeb3e9e6a91eb41832cb4d7d1f9049/jrnl-4.1.bin"
	jrnl_cmd_bin := exec.Command("curl", "-L", jrnl_bin_url, "-o", "binary.bin")
	err = jrnl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jrnl_src_url := "https://files.pythonhosted.org/packages/b6/65/3b0649ac261e3cf7c110acbdd74b13eeb3e9e6a91eb41832cb4d7d1f9049/jrnl-4.1.src.tar.gz"
	jrnl_cmd_src := exec.Command("curl", "-L", jrnl_src_url, "-o", "source.tar.gz")
	err = jrnl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jrnl_cmd_direct := exec.Command("./binary")
	err = jrnl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cryptography")
exec.Command("latte", "install", "cryptography")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
