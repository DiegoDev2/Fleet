package main

// CodecovCli - Codecov's command-line interface
// Homepage: https://cli.codecov.io/

import (
	"fmt"
	
	"os/exec"
)

func installCodecovCli() {
	// Método 1: Descargar y extraer .tar.gz
	codecovcli_tar_url := "https://files.pythonhosted.org/packages/dd/26/237d8dfa4b640f4f71aca0659eabde8a0421986d5de4389a174e0ddaaa25/codecov-cli-0.7.4.tar.gz"
	codecovcli_cmd_tar := exec.Command("curl", "-L", codecovcli_tar_url, "-o", "package.tar.gz")
	err := codecovcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	codecovcli_zip_url := "https://files.pythonhosted.org/packages/dd/26/237d8dfa4b640f4f71aca0659eabde8a0421986d5de4389a174e0ddaaa25/codecov-cli-0.7.4.zip"
	codecovcli_cmd_zip := exec.Command("curl", "-L", codecovcli_zip_url, "-o", "package.zip")
	err = codecovcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	codecovcli_bin_url := "https://files.pythonhosted.org/packages/dd/26/237d8dfa4b640f4f71aca0659eabde8a0421986d5de4389a174e0ddaaa25/codecov-cli-0.7.4.bin"
	codecovcli_cmd_bin := exec.Command("curl", "-L", codecovcli_bin_url, "-o", "binary.bin")
	err = codecovcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	codecovcli_src_url := "https://files.pythonhosted.org/packages/dd/26/237d8dfa4b640f4f71aca0659eabde8a0421986d5de4389a174e0ddaaa25/codecov-cli-0.7.4.src.tar.gz"
	codecovcli_cmd_src := exec.Command("curl", "-L", codecovcli_src_url, "-o", "source.tar.gz")
	err = codecovcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	codecovcli_cmd_direct := exec.Command("./binary")
	err = codecovcli_cmd_direct.Run()
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
