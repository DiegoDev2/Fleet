package main

// Mdbook - Create modern online books from Markdown files
// Homepage: https://rust-lang.github.io/mdBook/

import (
	"fmt"
	
	"os/exec"
)

func installMdbook() {
	// Método 1: Descargar y extraer .tar.gz
	mdbook_tar_url := "https://github.com/rust-lang/mdBook/archive/refs/tags/v0.4.40.tar.gz"
	mdbook_cmd_tar := exec.Command("curl", "-L", mdbook_tar_url, "-o", "package.tar.gz")
	err := mdbook_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mdbook_zip_url := "https://github.com/rust-lang/mdBook/archive/refs/tags/v0.4.40.zip"
	mdbook_cmd_zip := exec.Command("curl", "-L", mdbook_zip_url, "-o", "package.zip")
	err = mdbook_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mdbook_bin_url := "https://github.com/rust-lang/mdBook/archive/refs/tags/v0.4.40.bin"
	mdbook_cmd_bin := exec.Command("curl", "-L", mdbook_bin_url, "-o", "binary.bin")
	err = mdbook_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mdbook_src_url := "https://github.com/rust-lang/mdBook/archive/refs/tags/v0.4.40.src.tar.gz"
	mdbook_cmd_src := exec.Command("curl", "-L", mdbook_src_url, "-o", "source.tar.gz")
	err = mdbook_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mdbook_cmd_direct := exec.Command("./binary")
	err = mdbook_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
