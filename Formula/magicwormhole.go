package main

// MagicWormhole - Securely transfers data between computers
// Homepage: https://github.com/magic-wormhole/magic-wormhole

import (
	"fmt"
	
	"os/exec"
)

func installMagicWormhole() {
	// Método 1: Descargar y extraer .tar.gz
	magicwormhole_tar_url := "https://files.pythonhosted.org/packages/9d/7b/9320062c8b11e58dccd079dd62e3525bf1dac7f2b441c885d64a7ca7f045/magic-wormhole-0.15.0.tar.gz"
	magicwormhole_cmd_tar := exec.Command("curl", "-L", magicwormhole_tar_url, "-o", "package.tar.gz")
	err := magicwormhole_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	magicwormhole_zip_url := "https://files.pythonhosted.org/packages/9d/7b/9320062c8b11e58dccd079dd62e3525bf1dac7f2b441c885d64a7ca7f045/magic-wormhole-0.15.0.zip"
	magicwormhole_cmd_zip := exec.Command("curl", "-L", magicwormhole_zip_url, "-o", "package.zip")
	err = magicwormhole_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	magicwormhole_bin_url := "https://files.pythonhosted.org/packages/9d/7b/9320062c8b11e58dccd079dd62e3525bf1dac7f2b441c885d64a7ca7f045/magic-wormhole-0.15.0.bin"
	magicwormhole_cmd_bin := exec.Command("curl", "-L", magicwormhole_bin_url, "-o", "binary.bin")
	err = magicwormhole_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	magicwormhole_src_url := "https://files.pythonhosted.org/packages/9d/7b/9320062c8b11e58dccd079dd62e3525bf1dac7f2b441c885d64a7ca7f045/magic-wormhole-0.15.0.src.tar.gz"
	magicwormhole_cmd_src := exec.Command("curl", "-L", magicwormhole_src_url, "-o", "source.tar.gz")
	err = magicwormhole_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	magicwormhole_cmd_direct := exec.Command("./binary")
	err = magicwormhole_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cryptography")
	exec.Command("latte", "install", "cryptography").Run()
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
