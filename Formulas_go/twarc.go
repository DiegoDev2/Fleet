package main

// Twarc - Command-line tool and Python library for archiving Twitter JSON
// Homepage: https://github.com/DocNow/twarc

import (
	"fmt"
	
	"os/exec"
)

func installTwarc() {
	// Método 1: Descargar y extraer .tar.gz
	twarc_tar_url := "https://files.pythonhosted.org/packages/8a/ed/ac80b24ece6ee552f6deb39be34f01491cff4018cca8c5602c901dc08ecf/twarc-2.14.0.tar.gz"
	twarc_cmd_tar := exec.Command("curl", "-L", twarc_tar_url, "-o", "package.tar.gz")
	err := twarc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	twarc_zip_url := "https://files.pythonhosted.org/packages/8a/ed/ac80b24ece6ee552f6deb39be34f01491cff4018cca8c5602c901dc08ecf/twarc-2.14.0.zip"
	twarc_cmd_zip := exec.Command("curl", "-L", twarc_zip_url, "-o", "package.zip")
	err = twarc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	twarc_bin_url := "https://files.pythonhosted.org/packages/8a/ed/ac80b24ece6ee552f6deb39be34f01491cff4018cca8c5602c901dc08ecf/twarc-2.14.0.bin"
	twarc_cmd_bin := exec.Command("curl", "-L", twarc_bin_url, "-o", "binary.bin")
	err = twarc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	twarc_src_url := "https://files.pythonhosted.org/packages/8a/ed/ac80b24ece6ee552f6deb39be34f01491cff4018cca8c5602c901dc08ecf/twarc-2.14.0.src.tar.gz"
	twarc_cmd_src := exec.Command("curl", "-L", twarc_src_url, "-o", "source.tar.gz")
	err = twarc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	twarc_cmd_direct := exec.Command("./binary")
	err = twarc_cmd_direct.Run()
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
