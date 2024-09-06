package main

// Vulture - Find dead Python code
// Homepage: https://github.com/jendrikseipp/vulture

import (
	"fmt"
	
	"os/exec"
)

func installVulture() {
	// Método 1: Descargar y extraer .tar.gz
	vulture_tar_url := "https://files.pythonhosted.org/packages/da/70/29f296be6353598dfbbdf994f5496e6bf0776be6811c8491611a31aa15da/vulture-2.11.tar.gz"
	vulture_cmd_tar := exec.Command("curl", "-L", vulture_tar_url, "-o", "package.tar.gz")
	err := vulture_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vulture_zip_url := "https://files.pythonhosted.org/packages/da/70/29f296be6353598dfbbdf994f5496e6bf0776be6811c8491611a31aa15da/vulture-2.11.zip"
	vulture_cmd_zip := exec.Command("curl", "-L", vulture_zip_url, "-o", "package.zip")
	err = vulture_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vulture_bin_url := "https://files.pythonhosted.org/packages/da/70/29f296be6353598dfbbdf994f5496e6bf0776be6811c8491611a31aa15da/vulture-2.11.bin"
	vulture_cmd_bin := exec.Command("curl", "-L", vulture_bin_url, "-o", "binary.bin")
	err = vulture_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vulture_src_url := "https://files.pythonhosted.org/packages/da/70/29f296be6353598dfbbdf994f5496e6bf0776be6811c8491611a31aa15da/vulture-2.11.src.tar.gz"
	vulture_cmd_src := exec.Command("curl", "-L", vulture_src_url, "-o", "source.tar.gz")
	err = vulture_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vulture_cmd_direct := exec.Command("./binary")
	err = vulture_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
