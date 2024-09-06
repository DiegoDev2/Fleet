package main

// Todoman - Simple CalDAV-based todo manager
// Homepage: https://todoman.readthedocs.io/

import (
	"fmt"
	
	"os/exec"
)

func installTodoman() {
	// Método 1: Descargar y extraer .tar.gz
	todoman_tar_url := "https://files.pythonhosted.org/packages/fd/60/dbd18038cfe5a795d2e427b3ae4112c340966ed2d3a70303a4d59d7313eb/todoman-4.4.0.tar.gz"
	todoman_cmd_tar := exec.Command("curl", "-L", todoman_tar_url, "-o", "package.tar.gz")
	err := todoman_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	todoman_zip_url := "https://files.pythonhosted.org/packages/fd/60/dbd18038cfe5a795d2e427b3ae4112c340966ed2d3a70303a4d59d7313eb/todoman-4.4.0.zip"
	todoman_cmd_zip := exec.Command("curl", "-L", todoman_zip_url, "-o", "package.zip")
	err = todoman_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	todoman_bin_url := "https://files.pythonhosted.org/packages/fd/60/dbd18038cfe5a795d2e427b3ae4112c340966ed2d3a70303a4d59d7313eb/todoman-4.4.0.bin"
	todoman_cmd_bin := exec.Command("curl", "-L", todoman_bin_url, "-o", "binary.bin")
	err = todoman_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	todoman_src_url := "https://files.pythonhosted.org/packages/fd/60/dbd18038cfe5a795d2e427b3ae4112c340966ed2d3a70303a4d59d7313eb/todoman-4.4.0.src.tar.gz"
	todoman_cmd_src := exec.Command("curl", "-L", todoman_src_url, "-o", "source.tar.gz")
	err = todoman_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	todoman_cmd_direct := exec.Command("./binary")
	err = todoman_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: jq")
exec.Command("latte", "install", "jq")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
