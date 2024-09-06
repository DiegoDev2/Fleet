package main

// Nyx - Command-line monitor for Tor
// Homepage: https://nyx.torproject.org/

import (
	"fmt"
	
	"os/exec"
)

func installNyx() {
	// Método 1: Descargar y extraer .tar.gz
	nyx_tar_url := "https://files.pythonhosted.org/packages/f4/da/68419425cb0f64f996e2150045c7043c2bb61f77b5928c2156c26a21db88/nyx-2.1.0.tar.gz"
	nyx_cmd_tar := exec.Command("curl", "-L", nyx_tar_url, "-o", "package.tar.gz")
	err := nyx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	nyx_zip_url := "https://files.pythonhosted.org/packages/f4/da/68419425cb0f64f996e2150045c7043c2bb61f77b5928c2156c26a21db88/nyx-2.1.0.zip"
	nyx_cmd_zip := exec.Command("curl", "-L", nyx_zip_url, "-o", "package.zip")
	err = nyx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	nyx_bin_url := "https://files.pythonhosted.org/packages/f4/da/68419425cb0f64f996e2150045c7043c2bb61f77b5928c2156c26a21db88/nyx-2.1.0.bin"
	nyx_cmd_bin := exec.Command("curl", "-L", nyx_bin_url, "-o", "binary.bin")
	err = nyx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	nyx_src_url := "https://files.pythonhosted.org/packages/f4/da/68419425cb0f64f996e2150045c7043c2bb61f77b5928c2156c26a21db88/nyx-2.1.0.src.tar.gz"
	nyx_cmd_src := exec.Command("curl", "-L", nyx_src_url, "-o", "source.tar.gz")
	err = nyx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	nyx_cmd_direct := exec.Command("./binary")
	err = nyx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
