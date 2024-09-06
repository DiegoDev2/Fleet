package main

// Genometools - Versatile open source genome analysis software
// Homepage: https://genometools.org/

import (
	"fmt"
	
	"os/exec"
)

func installGenometools() {
	// Método 1: Descargar y extraer .tar.gz
	genometools_tar_url := "https://github.com/genometools/genometools/archive/refs/tags/v1.6.5.tar.gz"
	genometools_cmd_tar := exec.Command("curl", "-L", genometools_tar_url, "-o", "package.tar.gz")
	err := genometools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	genometools_zip_url := "https://github.com/genometools/genometools/archive/refs/tags/v1.6.5.zip"
	genometools_cmd_zip := exec.Command("curl", "-L", genometools_zip_url, "-o", "package.zip")
	err = genometools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	genometools_bin_url := "https://github.com/genometools/genometools/archive/refs/tags/v1.6.5.bin"
	genometools_cmd_bin := exec.Command("curl", "-L", genometools_bin_url, "-o", "binary.bin")
	err = genometools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	genometools_src_url := "https://github.com/genometools/genometools/archive/refs/tags/v1.6.5.src.tar.gz"
	genometools_cmd_src := exec.Command("curl", "-L", genometools_src_url, "-o", "source.tar.gz")
	err = genometools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	genometools_cmd_direct := exec.Command("./binary")
	err = genometools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
}
