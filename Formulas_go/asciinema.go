package main

// Asciinema - Record and share terminal sessions
// Homepage: https://asciinema.org

import (
	"fmt"
	
	"os/exec"
)

func installAsciinema() {
	// Método 1: Descargar y extraer .tar.gz
	asciinema_tar_url := "https://files.pythonhosted.org/packages/f1/19/45b405438e90ad5b9618f3df62e9b3edaa2b115b530e60bd4b363465c704/asciinema-2.4.0.tar.gz"
	asciinema_cmd_tar := exec.Command("curl", "-L", asciinema_tar_url, "-o", "package.tar.gz")
	err := asciinema_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	asciinema_zip_url := "https://files.pythonhosted.org/packages/f1/19/45b405438e90ad5b9618f3df62e9b3edaa2b115b530e60bd4b363465c704/asciinema-2.4.0.zip"
	asciinema_cmd_zip := exec.Command("curl", "-L", asciinema_zip_url, "-o", "package.zip")
	err = asciinema_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	asciinema_bin_url := "https://files.pythonhosted.org/packages/f1/19/45b405438e90ad5b9618f3df62e9b3edaa2b115b530e60bd4b363465c704/asciinema-2.4.0.bin"
	asciinema_cmd_bin := exec.Command("curl", "-L", asciinema_bin_url, "-o", "binary.bin")
	err = asciinema_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	asciinema_src_url := "https://files.pythonhosted.org/packages/f1/19/45b405438e90ad5b9618f3df62e9b3edaa2b115b530e60bd4b363465c704/asciinema-2.4.0.src.tar.gz"
	asciinema_cmd_src := exec.Command("curl", "-L", asciinema_src_url, "-o", "source.tar.gz")
	err = asciinema_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	asciinema_cmd_direct := exec.Command("./binary")
	err = asciinema_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
