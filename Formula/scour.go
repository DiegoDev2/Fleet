package main

// Scour - SVG file scrubber
// Homepage: https://www.codedread.com/scour/

import (
	"fmt"
	
	"os/exec"
)

func installScour() {
	// Método 1: Descargar y extraer .tar.gz
	scour_tar_url := "https://files.pythonhosted.org/packages/75/19/f519ef8aa2f379935a44212c5744e2b3a46173bf04e0110fb7f4af4028c9/scour-0.38.2.tar.gz"
	scour_cmd_tar := exec.Command("curl", "-L", scour_tar_url, "-o", "package.tar.gz")
	err := scour_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	scour_zip_url := "https://files.pythonhosted.org/packages/75/19/f519ef8aa2f379935a44212c5744e2b3a46173bf04e0110fb7f4af4028c9/scour-0.38.2.zip"
	scour_cmd_zip := exec.Command("curl", "-L", scour_zip_url, "-o", "package.zip")
	err = scour_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	scour_bin_url := "https://files.pythonhosted.org/packages/75/19/f519ef8aa2f379935a44212c5744e2b3a46173bf04e0110fb7f4af4028c9/scour-0.38.2.bin"
	scour_cmd_bin := exec.Command("curl", "-L", scour_bin_url, "-o", "binary.bin")
	err = scour_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	scour_src_url := "https://files.pythonhosted.org/packages/75/19/f519ef8aa2f379935a44212c5744e2b3a46173bf04e0110fb7f4af4028c9/scour-0.38.2.src.tar.gz"
	scour_cmd_src := exec.Command("curl", "-L", scour_src_url, "-o", "source.tar.gz")
	err = scour_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	scour_cmd_direct := exec.Command("./binary")
	err = scour_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
