package main

// Codelimit - Your Refactoring Alarm
// Homepage: https://github.com/getcodelimit/codelimit

import (
	"fmt"
	
	"os/exec"
)

func installCodelimit() {
	// Método 1: Descargar y extraer .tar.gz
	codelimit_tar_url := "https://files.pythonhosted.org/packages/9b/54/f6fe026726846c0504da0f641e00738c4dbb2ba527dc642344186571fda8/codelimit-0.9.5.tar.gz"
	codelimit_cmd_tar := exec.Command("curl", "-L", codelimit_tar_url, "-o", "package.tar.gz")
	err := codelimit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	codelimit_zip_url := "https://files.pythonhosted.org/packages/9b/54/f6fe026726846c0504da0f641e00738c4dbb2ba527dc642344186571fda8/codelimit-0.9.5.zip"
	codelimit_cmd_zip := exec.Command("curl", "-L", codelimit_zip_url, "-o", "package.zip")
	err = codelimit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	codelimit_bin_url := "https://files.pythonhosted.org/packages/9b/54/f6fe026726846c0504da0f641e00738c4dbb2ba527dc642344186571fda8/codelimit-0.9.5.bin"
	codelimit_cmd_bin := exec.Command("curl", "-L", codelimit_bin_url, "-o", "binary.bin")
	err = codelimit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	codelimit_src_url := "https://files.pythonhosted.org/packages/9b/54/f6fe026726846c0504da0f641e00738c4dbb2ba527dc642344186571fda8/codelimit-0.9.5.src.tar.gz"
	codelimit_cmd_src := exec.Command("curl", "-L", codelimit_src_url, "-o", "source.tar.gz")
	err = codelimit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	codelimit_cmd_direct := exec.Command("./binary")
	err = codelimit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
