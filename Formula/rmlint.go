package main

// Rmlint - Extremely fast tool to remove dupes and other lint from your filesystem
// Homepage: https://github.com/sahib/rmlint

import (
	"fmt"
	
	"os/exec"
)

func installRmlint() {
	// Método 1: Descargar y extraer .tar.gz
	rmlint_tar_url := "https://github.com/sahib/rmlint/archive/refs/tags/v2.10.2.tar.gz"
	rmlint_cmd_tar := exec.Command("curl", "-L", rmlint_tar_url, "-o", "package.tar.gz")
	err := rmlint_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rmlint_zip_url := "https://github.com/sahib/rmlint/archive/refs/tags/v2.10.2.zip"
	rmlint_cmd_zip := exec.Command("curl", "-L", rmlint_zip_url, "-o", "package.zip")
	err = rmlint_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rmlint_bin_url := "https://github.com/sahib/rmlint/archive/refs/tags/v2.10.2.bin"
	rmlint_cmd_bin := exec.Command("curl", "-L", rmlint_bin_url, "-o", "binary.bin")
	err = rmlint_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rmlint_src_url := "https://github.com/sahib/rmlint/archive/refs/tags/v2.10.2.src.tar.gz"
	rmlint_cmd_src := exec.Command("curl", "-L", rmlint_src_url, "-o", "source.tar.gz")
	err = rmlint_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rmlint_cmd_direct := exec.Command("./binary")
	err = rmlint_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: scons")
	exec.Command("latte", "install", "scons").Run()
	fmt.Println("Instalando dependencia: sphinx-doc")
	exec.Command("latte", "install", "sphinx-doc").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: json-glib")
	exec.Command("latte", "install", "json-glib").Run()
	fmt.Println("Instalando dependencia: elfutils")
	exec.Command("latte", "install", "elfutils").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
