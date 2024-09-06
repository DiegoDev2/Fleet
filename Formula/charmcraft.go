package main

// Charmcraft - Tool to build charms and publish them on Charmhub
// Homepage: https://charmhub.io

import (
	"fmt"
	
	"os/exec"
)

func installCharmcraft() {
	// Método 1: Descargar y extraer .tar.gz
	charmcraft_tar_url := "https://files.pythonhosted.org/packages/09/f3/bfa495c22cf4fb59d541feb4f931c10895d194f12160a2c9c29165ed0da1/charmcraft-3.2.0.tar.gz"
	charmcraft_cmd_tar := exec.Command("curl", "-L", charmcraft_tar_url, "-o", "package.tar.gz")
	err := charmcraft_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	charmcraft_zip_url := "https://files.pythonhosted.org/packages/09/f3/bfa495c22cf4fb59d541feb4f931c10895d194f12160a2c9c29165ed0da1/charmcraft-3.2.0.zip"
	charmcraft_cmd_zip := exec.Command("curl", "-L", charmcraft_zip_url, "-o", "package.zip")
	err = charmcraft_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	charmcraft_bin_url := "https://files.pythonhosted.org/packages/09/f3/bfa495c22cf4fb59d541feb4f931c10895d194f12160a2c9c29165ed0da1/charmcraft-3.2.0.bin"
	charmcraft_cmd_bin := exec.Command("curl", "-L", charmcraft_bin_url, "-o", "binary.bin")
	err = charmcraft_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	charmcraft_src_url := "https://files.pythonhosted.org/packages/09/f3/bfa495c22cf4fb59d541feb4f931c10895d194f12160a2c9c29165ed0da1/charmcraft-3.2.0.src.tar.gz"
	charmcraft_cmd_src := exec.Command("curl", "-L", charmcraft_src_url, "-o", "source.tar.gz")
	err = charmcraft_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	charmcraft_cmd_direct := exec.Command("./binary")
	err = charmcraft_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
	exec.Command("latte", "install", "rust").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: pygit2")
	exec.Command("latte", "install", "pygit2").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
