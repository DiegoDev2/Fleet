package main

// Instalooter - Download any picture or video associated from an Instagram profile
// Homepage: https://github.com/althonos/instalooter

import (
	"fmt"
	
	"os/exec"
)

func installInstalooter() {
	// Método 1: Descargar y extraer .tar.gz
	instalooter_tar_url := "https://files.pythonhosted.org/packages/30/13/907e6aaba6280e1001080ab47e750068ffc5fb7174203985b3c9d678e3f2/instalooter-2.4.4.tar.gz"
	instalooter_cmd_tar := exec.Command("curl", "-L", instalooter_tar_url, "-o", "package.tar.gz")
	err := instalooter_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	instalooter_zip_url := "https://files.pythonhosted.org/packages/30/13/907e6aaba6280e1001080ab47e750068ffc5fb7174203985b3c9d678e3f2/instalooter-2.4.4.zip"
	instalooter_cmd_zip := exec.Command("curl", "-L", instalooter_zip_url, "-o", "package.zip")
	err = instalooter_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	instalooter_bin_url := "https://files.pythonhosted.org/packages/30/13/907e6aaba6280e1001080ab47e750068ffc5fb7174203985b3c9d678e3f2/instalooter-2.4.4.bin"
	instalooter_cmd_bin := exec.Command("curl", "-L", instalooter_bin_url, "-o", "binary.bin")
	err = instalooter_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	instalooter_src_url := "https://files.pythonhosted.org/packages/30/13/907e6aaba6280e1001080ab47e750068ffc5fb7174203985b3c9d678e3f2/instalooter-2.4.4.src.tar.gz"
	instalooter_cmd_src := exec.Command("curl", "-L", instalooter_src_url, "-o", "source.tar.gz")
	err = instalooter_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	instalooter_cmd_direct := exec.Command("./binary")
	err = instalooter_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
