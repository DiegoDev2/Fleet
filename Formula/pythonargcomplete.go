package main

// PythonArgcomplete - Tab completion for Python argparse
// Homepage: https://kislyuk.github.io/argcomplete/

import (
	"fmt"
	
	"os/exec"
)

func installPythonArgcomplete() {
	// Método 1: Descargar y extraer .tar.gz
	pythonargcomplete_tar_url := "https://files.pythonhosted.org/packages/75/33/a3d23a2e9ac78f9eaf1fce7490fee430d43ca7d42c65adabbb36a2b28ff6/argcomplete-3.5.0.tar.gz"
	pythonargcomplete_cmd_tar := exec.Command("curl", "-L", pythonargcomplete_tar_url, "-o", "package.tar.gz")
	err := pythonargcomplete_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pythonargcomplete_zip_url := "https://files.pythonhosted.org/packages/75/33/a3d23a2e9ac78f9eaf1fce7490fee430d43ca7d42c65adabbb36a2b28ff6/argcomplete-3.5.0.zip"
	pythonargcomplete_cmd_zip := exec.Command("curl", "-L", pythonargcomplete_zip_url, "-o", "package.zip")
	err = pythonargcomplete_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pythonargcomplete_bin_url := "https://files.pythonhosted.org/packages/75/33/a3d23a2e9ac78f9eaf1fce7490fee430d43ca7d42c65adabbb36a2b28ff6/argcomplete-3.5.0.bin"
	pythonargcomplete_cmd_bin := exec.Command("curl", "-L", pythonargcomplete_bin_url, "-o", "binary.bin")
	err = pythonargcomplete_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pythonargcomplete_src_url := "https://files.pythonhosted.org/packages/75/33/a3d23a2e9ac78f9eaf1fce7490fee430d43ca7d42c65adabbb36a2b28ff6/argcomplete-3.5.0.src.tar.gz"
	pythonargcomplete_cmd_src := exec.Command("curl", "-L", pythonargcomplete_src_url, "-o", "source.tar.gz")
	err = pythonargcomplete_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pythonargcomplete_cmd_direct := exec.Command("./binary")
	err = pythonargcomplete_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
