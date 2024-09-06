package main

// Johnnydep - Display dependency tree of Python distribution
// Homepage: https://github.com/wimglenn/johnnydep

import (
	"fmt"
	
	"os/exec"
)

func installJohnnydep() {
	// Método 1: Descargar y extraer .tar.gz
	johnnydep_tar_url := "https://files.pythonhosted.org/packages/e5/57/ccdd7ab46a4a06fae442555fe90a02d551009d765b79b99a942b2df330c5/johnnydep-1.20.5.tar.gz"
	johnnydep_cmd_tar := exec.Command("curl", "-L", johnnydep_tar_url, "-o", "package.tar.gz")
	err := johnnydep_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	johnnydep_zip_url := "https://files.pythonhosted.org/packages/e5/57/ccdd7ab46a4a06fae442555fe90a02d551009d765b79b99a942b2df330c5/johnnydep-1.20.5.zip"
	johnnydep_cmd_zip := exec.Command("curl", "-L", johnnydep_zip_url, "-o", "package.zip")
	err = johnnydep_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	johnnydep_bin_url := "https://files.pythonhosted.org/packages/e5/57/ccdd7ab46a4a06fae442555fe90a02d551009d765b79b99a942b2df330c5/johnnydep-1.20.5.bin"
	johnnydep_cmd_bin := exec.Command("curl", "-L", johnnydep_bin_url, "-o", "binary.bin")
	err = johnnydep_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	johnnydep_src_url := "https://files.pythonhosted.org/packages/e5/57/ccdd7ab46a4a06fae442555fe90a02d551009d765b79b99a942b2df330c5/johnnydep-1.20.5.src.tar.gz"
	johnnydep_cmd_src := exec.Command("curl", "-L", johnnydep_src_url, "-o", "source.tar.gz")
	err = johnnydep_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	johnnydep_cmd_direct := exec.Command("./binary")
	err = johnnydep_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
