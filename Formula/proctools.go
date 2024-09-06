package main

// Proctools - OpenBSD and Darwin versions of pgrep, pkill, and pfind
// Homepage: https://proctools.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installProctools() {
	// Método 1: Descargar y extraer .tar.gz
	proctools_tar_url := "https://downloads.sourceforge.net/project/proctools/proctools/0.4pre1/proctools-0.4pre1.tar.gz"
	proctools_cmd_tar := exec.Command("curl", "-L", proctools_tar_url, "-o", "package.tar.gz")
	err := proctools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	proctools_zip_url := "https://downloads.sourceforge.net/project/proctools/proctools/0.4pre1/proctools-0.4pre1.zip"
	proctools_cmd_zip := exec.Command("curl", "-L", proctools_zip_url, "-o", "package.zip")
	err = proctools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	proctools_bin_url := "https://downloads.sourceforge.net/project/proctools/proctools/0.4pre1/proctools-0.4pre1.bin"
	proctools_cmd_bin := exec.Command("curl", "-L", proctools_bin_url, "-o", "binary.bin")
	err = proctools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	proctools_src_url := "https://downloads.sourceforge.net/project/proctools/proctools/0.4pre1/proctools-0.4pre1.src.tar.gz"
	proctools_cmd_src := exec.Command("curl", "-L", proctools_src_url, "-o", "source.tar.gz")
	err = proctools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	proctools_cmd_direct := exec.Command("./binary")
	err = proctools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bsdmake")
	exec.Command("latte", "install", "bsdmake").Run()
}
