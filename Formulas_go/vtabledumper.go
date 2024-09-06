package main

// VtableDumper - List contents of virtual tables in a shared library
// Homepage: https://github.com/lvc/vtable-dumper

import (
	"fmt"
	
	"os/exec"
)

func installVtableDumper() {
	// Método 1: Descargar y extraer .tar.gz
	vtabledumper_tar_url := "https://github.com/lvc/vtable-dumper/archive/refs/tags/1.2.tar.gz"
	vtabledumper_cmd_tar := exec.Command("curl", "-L", vtabledumper_tar_url, "-o", "package.tar.gz")
	err := vtabledumper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vtabledumper_zip_url := "https://github.com/lvc/vtable-dumper/archive/refs/tags/1.2.zip"
	vtabledumper_cmd_zip := exec.Command("curl", "-L", vtabledumper_zip_url, "-o", "package.zip")
	err = vtabledumper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vtabledumper_bin_url := "https://github.com/lvc/vtable-dumper/archive/refs/tags/1.2.bin"
	vtabledumper_cmd_bin := exec.Command("curl", "-L", vtabledumper_bin_url, "-o", "binary.bin")
	err = vtabledumper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vtabledumper_src_url := "https://github.com/lvc/vtable-dumper/archive/refs/tags/1.2.src.tar.gz"
	vtabledumper_cmd_src := exec.Command("curl", "-L", vtabledumper_src_url, "-o", "source.tar.gz")
	err = vtabledumper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vtabledumper_cmd_direct := exec.Command("./binary")
	err = vtabledumper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: elfutils")
exec.Command("latte", "install", "elfutils")
}
