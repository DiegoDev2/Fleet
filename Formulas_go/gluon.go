package main

// Gluon - Static, type inferred and embeddable language written in Rust
// Homepage: https://gluon-lang.org

import (
	"fmt"
	
	"os/exec"
)

func installGluon() {
	// Método 1: Descargar y extraer .tar.gz
	gluon_tar_url := "https://github.com/gluon-lang/gluon/archive/refs/tags/v0.18.2.tar.gz"
	gluon_cmd_tar := exec.Command("curl", "-L", gluon_tar_url, "-o", "package.tar.gz")
	err := gluon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gluon_zip_url := "https://github.com/gluon-lang/gluon/archive/refs/tags/v0.18.2.zip"
	gluon_cmd_zip := exec.Command("curl", "-L", gluon_zip_url, "-o", "package.zip")
	err = gluon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gluon_bin_url := "https://github.com/gluon-lang/gluon/archive/refs/tags/v0.18.2.bin"
	gluon_cmd_bin := exec.Command("curl", "-L", gluon_bin_url, "-o", "binary.bin")
	err = gluon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gluon_src_url := "https://github.com/gluon-lang/gluon/archive/refs/tags/v0.18.2.src.tar.gz"
	gluon_cmd_src := exec.Command("curl", "-L", gluon_src_url, "-o", "source.tar.gz")
	err = gluon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gluon_cmd_direct := exec.Command("./binary")
	err = gluon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
