package main

// Remake - GNU Make with improved error handling, tracing, and a debugger
// Homepage: https://bashdb.sourceforge.net/remake/

import (
	"fmt"
	
	"os/exec"
)

func installRemake() {
	// Método 1: Descargar y extraer .tar.gz
	remake_tar_url := "https://downloads.sourceforge.net/project/bashdb/remake/4.3%2Bdbg-1.6/remake-4.3%2Bdbg-1.6.tar.gz"
	remake_cmd_tar := exec.Command("curl", "-L", remake_tar_url, "-o", "package.tar.gz")
	err := remake_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	remake_zip_url := "https://downloads.sourceforge.net/project/bashdb/remake/4.3%2Bdbg-1.6/remake-4.3%2Bdbg-1.6.zip"
	remake_cmd_zip := exec.Command("curl", "-L", remake_zip_url, "-o", "package.zip")
	err = remake_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	remake_bin_url := "https://downloads.sourceforge.net/project/bashdb/remake/4.3%2Bdbg-1.6/remake-4.3%2Bdbg-1.6.bin"
	remake_cmd_bin := exec.Command("curl", "-L", remake_bin_url, "-o", "binary.bin")
	err = remake_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	remake_src_url := "https://downloads.sourceforge.net/project/bashdb/remake/4.3%2Bdbg-1.6/remake-4.3%2Bdbg-1.6.src.tar.gz"
	remake_cmd_src := exec.Command("curl", "-L", remake_src_url, "-o", "source.tar.gz")
	err = remake_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	remake_cmd_direct := exec.Command("./binary")
	err = remake_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: readline")
	exec.Command("latte", "install", "readline").Run()
}
