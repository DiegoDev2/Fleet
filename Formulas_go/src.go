package main

// Src - Simple revision control: RCS reloaded with a modern UI
// Homepage: http://www.catb.org/~esr/src/

import (
	"fmt"
	
	"os/exec"
)

func installSrc() {
	// Método 1: Descargar y extraer .tar.gz
	src_tar_url := "http://www.catb.org/~esr/src/src-1.40.tar.gz"
	src_cmd_tar := exec.Command("curl", "-L", src_tar_url, "-o", "package.tar.gz")
	err := src_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	src_zip_url := "http://www.catb.org/~esr/src/src-1.40.zip"
	src_cmd_zip := exec.Command("curl", "-L", src_zip_url, "-o", "package.zip")
	err = src_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	src_bin_url := "http://www.catb.org/~esr/src/src-1.40.bin"
	src_cmd_bin := exec.Command("curl", "-L", src_bin_url, "-o", "binary.bin")
	err = src_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	src_src_url := "http://www.catb.org/~esr/src/src-1.40.src.tar.gz"
	src_cmd_src := exec.Command("curl", "-L", src_src_url, "-o", "source.tar.gz")
	err = src_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	src_cmd_direct := exec.Command("./binary")
	err = src_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoc")
exec.Command("latte", "install", "asciidoc")
	fmt.Println("Instalando dependencia: rcs")
exec.Command("latte", "install", "rcs")
}
