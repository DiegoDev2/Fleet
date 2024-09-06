package main

// Brag - Download and assemble multipart binaries from newsgroups
// Homepage: https://brag.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installBrag() {
	// Método 1: Descargar y extraer .tar.gz
	brag_tar_url := "https://downloads.sourceforge.net/project/brag/brag/1.4.3/brag-1.4.3.tar.gz"
	brag_cmd_tar := exec.Command("curl", "-L", brag_tar_url, "-o", "package.tar.gz")
	err := brag_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	brag_zip_url := "https://downloads.sourceforge.net/project/brag/brag/1.4.3/brag-1.4.3.zip"
	brag_cmd_zip := exec.Command("curl", "-L", brag_zip_url, "-o", "package.zip")
	err = brag_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	brag_bin_url := "https://downloads.sourceforge.net/project/brag/brag/1.4.3/brag-1.4.3.bin"
	brag_cmd_bin := exec.Command("curl", "-L", brag_bin_url, "-o", "binary.bin")
	err = brag_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	brag_src_url := "https://downloads.sourceforge.net/project/brag/brag/1.4.3/brag-1.4.3.src.tar.gz"
	brag_cmd_src := exec.Command("curl", "-L", brag_src_url, "-o", "source.tar.gz")
	err = brag_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	brag_cmd_direct := exec.Command("./binary")
	err = brag_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: uudeview")
exec.Command("latte", "install", "uudeview")
	fmt.Println("Instalando dependencia: tcl-tk")
exec.Command("latte", "install", "tcl-tk")
}
