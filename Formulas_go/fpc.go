package main

// Fpc - Free Pascal: multi-architecture Pascal compiler
// Homepage: https://www.freepascal.org/

import (
	"fmt"
	
	"os/exec"
)

func installFpc() {
	// Método 1: Descargar y extraer .tar.gz
	fpc_tar_url := "https://downloads.sourceforge.net/project/freepascal/Source/3.2.2/fpc-3.2.2.source.tar.gz"
	fpc_cmd_tar := exec.Command("curl", "-L", fpc_tar_url, "-o", "package.tar.gz")
	err := fpc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fpc_zip_url := "https://downloads.sourceforge.net/project/freepascal/Source/3.2.2/fpc-3.2.2.source.zip"
	fpc_cmd_zip := exec.Command("curl", "-L", fpc_zip_url, "-o", "package.zip")
	err = fpc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fpc_bin_url := "https://downloads.sourceforge.net/project/freepascal/Source/3.2.2/fpc-3.2.2.source.bin"
	fpc_cmd_bin := exec.Command("curl", "-L", fpc_bin_url, "-o", "binary.bin")
	err = fpc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fpc_src_url := "https://downloads.sourceforge.net/project/freepascal/Source/3.2.2/fpc-3.2.2.source.src.tar.gz"
	fpc_cmd_src := exec.Command("curl", "-L", fpc_src_url, "-o", "source.tar.gz")
	err = fpc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fpc_cmd_direct := exec.Command("./binary")
	err = fpc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mesa")
exec.Command("latte", "install", "mesa")
}
