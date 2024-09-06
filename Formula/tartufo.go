package main

// Tartufo - Searches through git repositories for high entropy strings and secrets
// Homepage: https://tartufo.readthedocs.io/en/stable/

import (
	"fmt"
	
	"os/exec"
)

func installTartufo() {
	// Método 1: Descargar y extraer .tar.gz
	tartufo_tar_url := "https://files.pythonhosted.org/packages/f3/be/a004a02e3b2be08c998f66f391df238de701320af3f0a0438e724db943e2/tartufo-5.0.1.tar.gz"
	tartufo_cmd_tar := exec.Command("curl", "-L", tartufo_tar_url, "-o", "package.tar.gz")
	err := tartufo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tartufo_zip_url := "https://files.pythonhosted.org/packages/f3/be/a004a02e3b2be08c998f66f391df238de701320af3f0a0438e724db943e2/tartufo-5.0.1.zip"
	tartufo_cmd_zip := exec.Command("curl", "-L", tartufo_zip_url, "-o", "package.zip")
	err = tartufo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tartufo_bin_url := "https://files.pythonhosted.org/packages/f3/be/a004a02e3b2be08c998f66f391df238de701320af3f0a0438e724db943e2/tartufo-5.0.1.bin"
	tartufo_cmd_bin := exec.Command("curl", "-L", tartufo_bin_url, "-o", "binary.bin")
	err = tartufo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tartufo_src_url := "https://files.pythonhosted.org/packages/f3/be/a004a02e3b2be08c998f66f391df238de701320af3f0a0438e724db943e2/tartufo-5.0.1.src.tar.gz"
	tartufo_cmd_src := exec.Command("curl", "-L", tartufo_src_url, "-o", "source.tar.gz")
	err = tartufo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tartufo_cmd_direct := exec.Command("./binary")
	err = tartufo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pygit2")
	exec.Command("latte", "install", "pygit2").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
