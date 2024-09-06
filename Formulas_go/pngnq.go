package main

// Pngnq - Tool for optimizing PNG images
// Homepage: https://pngnq.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installPngnq() {
	// Método 1: Descargar y extraer .tar.gz
	pngnq_tar_url := "https://downloads.sourceforge.net/project/pngnq/pngnq/1.1/pngnq-1.1.tar.gz"
	pngnq_cmd_tar := exec.Command("curl", "-L", pngnq_tar_url, "-o", "package.tar.gz")
	err := pngnq_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pngnq_zip_url := "https://downloads.sourceforge.net/project/pngnq/pngnq/1.1/pngnq-1.1.zip"
	pngnq_cmd_zip := exec.Command("curl", "-L", pngnq_zip_url, "-o", "package.zip")
	err = pngnq_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pngnq_bin_url := "https://downloads.sourceforge.net/project/pngnq/pngnq/1.1/pngnq-1.1.bin"
	pngnq_cmd_bin := exec.Command("curl", "-L", pngnq_bin_url, "-o", "binary.bin")
	err = pngnq_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pngnq_src_url := "https://downloads.sourceforge.net/project/pngnq/pngnq/1.1/pngnq-1.1.src.tar.gz"
	pngnq_cmd_src := exec.Command("curl", "-L", pngnq_src_url, "-o", "source.tar.gz")
	err = pngnq_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pngnq_cmd_direct := exec.Command("./binary")
	err = pngnq_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
}
