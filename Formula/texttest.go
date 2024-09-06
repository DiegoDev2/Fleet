package main

// Texttest - Tool for text-based Approval Testing
// Homepage: https://www.texttest.org/

import (
	"fmt"
	
	"os/exec"
)

func installTexttest() {
	// Método 1: Descargar y extraer .tar.gz
	texttest_tar_url := "https://files.pythonhosted.org/packages/7b/14/e52c96906f1d397c776c4940f68e9b44cae6b1a1aaba915c372638c3b48f/TextTest-4.3.1.tar.gz"
	texttest_cmd_tar := exec.Command("curl", "-L", texttest_tar_url, "-o", "package.tar.gz")
	err := texttest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	texttest_zip_url := "https://files.pythonhosted.org/packages/7b/14/e52c96906f1d397c776c4940f68e9b44cae6b1a1aaba915c372638c3b48f/TextTest-4.3.1.zip"
	texttest_cmd_zip := exec.Command("curl", "-L", texttest_zip_url, "-o", "package.zip")
	err = texttest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	texttest_bin_url := "https://files.pythonhosted.org/packages/7b/14/e52c96906f1d397c776c4940f68e9b44cae6b1a1aaba915c372638c3b48f/TextTest-4.3.1.bin"
	texttest_cmd_bin := exec.Command("curl", "-L", texttest_bin_url, "-o", "binary.bin")
	err = texttest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	texttest_src_url := "https://files.pythonhosted.org/packages/7b/14/e52c96906f1d397c776c4940f68e9b44cae6b1a1aaba915c372638c3b48f/TextTest-4.3.1.src.tar.gz"
	texttest_cmd_src := exec.Command("curl", "-L", texttest_src_url, "-o", "source.tar.gz")
	err = texttest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	texttest_cmd_direct := exec.Command("./binary")
	err = texttest_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: adwaita-icon-theme")
	exec.Command("latte", "install", "adwaita-icon-theme").Run()
	fmt.Println("Instalando dependencia: gtk+3")
	exec.Command("latte", "install", "gtk+3").Run()
	fmt.Println("Instalando dependencia: pygobject3")
	exec.Command("latte", "install", "pygobject3").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
