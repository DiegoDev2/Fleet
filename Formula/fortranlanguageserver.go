package main

// FortranLanguageServer - Language Server for Fortran
// Homepage: https://github.com/hansec/fortran-language-server

import (
	"fmt"
	
	"os/exec"
)

func installFortranLanguageServer() {
	// Método 1: Descargar y extraer .tar.gz
	fortranlanguageserver_tar_url := "https://files.pythonhosted.org/packages/72/46/eb2c733e920a33409906aa145bde93b015f7f77c9bb8bdf65faa8c823998/fortran-language-server-1.12.0.tar.gz"
	fortranlanguageserver_cmd_tar := exec.Command("curl", "-L", fortranlanguageserver_tar_url, "-o", "package.tar.gz")
	err := fortranlanguageserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fortranlanguageserver_zip_url := "https://files.pythonhosted.org/packages/72/46/eb2c733e920a33409906aa145bde93b015f7f77c9bb8bdf65faa8c823998/fortran-language-server-1.12.0.zip"
	fortranlanguageserver_cmd_zip := exec.Command("curl", "-L", fortranlanguageserver_zip_url, "-o", "package.zip")
	err = fortranlanguageserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fortranlanguageserver_bin_url := "https://files.pythonhosted.org/packages/72/46/eb2c733e920a33409906aa145bde93b015f7f77c9bb8bdf65faa8c823998/fortran-language-server-1.12.0.bin"
	fortranlanguageserver_cmd_bin := exec.Command("curl", "-L", fortranlanguageserver_bin_url, "-o", "binary.bin")
	err = fortranlanguageserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fortranlanguageserver_src_url := "https://files.pythonhosted.org/packages/72/46/eb2c733e920a33409906aa145bde93b015f7f77c9bb8bdf65faa8c823998/fortran-language-server-1.12.0.src.tar.gz"
	fortranlanguageserver_cmd_src := exec.Command("curl", "-L", fortranlanguageserver_src_url, "-o", "source.tar.gz")
	err = fortranlanguageserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fortranlanguageserver_cmd_direct := exec.Command("./binary")
	err = fortranlanguageserver_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
