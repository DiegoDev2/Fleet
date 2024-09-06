package main

// Elektra - Framework to access config settings in a global key database
// Homepage: https://www.libelektra.org

import (
	"fmt"
	
	"os/exec"
)

func installElektra() {
	// Método 1: Descargar y extraer .tar.gz
	elektra_tar_url := "https://www.libelektra.org/ftp/elektra/releases/elektra-0.11.0.tar.gz"
	elektra_cmd_tar := exec.Command("curl", "-L", elektra_tar_url, "-o", "package.tar.gz")
	err := elektra_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	elektra_zip_url := "https://www.libelektra.org/ftp/elektra/releases/elektra-0.11.0.zip"
	elektra_cmd_zip := exec.Command("curl", "-L", elektra_zip_url, "-o", "package.zip")
	err = elektra_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	elektra_bin_url := "https://www.libelektra.org/ftp/elektra/releases/elektra-0.11.0.bin"
	elektra_cmd_bin := exec.Command("curl", "-L", elektra_bin_url, "-o", "binary.bin")
	err = elektra_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	elektra_src_url := "https://www.libelektra.org/ftp/elektra/releases/elektra-0.11.0.src.tar.gz"
	elektra_cmd_src := exec.Command("curl", "-L", elektra_src_url, "-o", "source.tar.gz")
	err = elektra_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	elektra_cmd_direct := exec.Command("./binary")
	err = elektra_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: doxygen")
exec.Command("latte", "install", "doxygen")
}
