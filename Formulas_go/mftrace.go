package main

// Mftrace - Trace TeX bitmap font to PFA, PFB, or TTF font
// Homepage: https://lilypond.org/mftrace/

import (
	"fmt"
	
	"os/exec"
)

func installMftrace() {
	// Método 1: Descargar y extraer .tar.gz
	mftrace_tar_url := "https://lilypond.org/downloads/sources/mftrace/mftrace-1.2.20.tar.gz"
	mftrace_cmd_tar := exec.Command("curl", "-L", mftrace_tar_url, "-o", "package.tar.gz")
	err := mftrace_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mftrace_zip_url := "https://lilypond.org/downloads/sources/mftrace/mftrace-1.2.20.zip"
	mftrace_cmd_zip := exec.Command("curl", "-L", mftrace_zip_url, "-o", "package.zip")
	err = mftrace_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mftrace_bin_url := "https://lilypond.org/downloads/sources/mftrace/mftrace-1.2.20.bin"
	mftrace_cmd_bin := exec.Command("curl", "-L", mftrace_bin_url, "-o", "binary.bin")
	err = mftrace_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mftrace_src_url := "https://lilypond.org/downloads/sources/mftrace/mftrace-1.2.20.src.tar.gz"
	mftrace_cmd_src := exec.Command("curl", "-L", mftrace_src_url, "-o", "source.tar.gz")
	err = mftrace_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mftrace_cmd_direct := exec.Command("./binary")
	err = mftrace_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: fontforge")
exec.Command("latte", "install", "fontforge")
	fmt.Println("Instalando dependencia: potrace")
exec.Command("latte", "install", "potrace")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: t1utils")
exec.Command("latte", "install", "t1utils")
}
