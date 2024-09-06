package main

// Cmuclmtk - Language model tools (from CMU Sphinx)
// Homepage: https://cmusphinx.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installCmuclmtk() {
	// Método 1: Descargar y extraer .tar.gz
	cmuclmtk_tar_url := "https://downloads.sourceforge.net/project/cmusphinx/cmuclmtk/0.7/cmuclmtk-0.7.tar.gz"
	cmuclmtk_cmd_tar := exec.Command("curl", "-L", cmuclmtk_tar_url, "-o", "package.tar.gz")
	err := cmuclmtk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmuclmtk_zip_url := "https://downloads.sourceforge.net/project/cmusphinx/cmuclmtk/0.7/cmuclmtk-0.7.zip"
	cmuclmtk_cmd_zip := exec.Command("curl", "-L", cmuclmtk_zip_url, "-o", "package.zip")
	err = cmuclmtk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmuclmtk_bin_url := "https://downloads.sourceforge.net/project/cmusphinx/cmuclmtk/0.7/cmuclmtk-0.7.bin"
	cmuclmtk_cmd_bin := exec.Command("curl", "-L", cmuclmtk_bin_url, "-o", "binary.bin")
	err = cmuclmtk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmuclmtk_src_url := "https://downloads.sourceforge.net/project/cmusphinx/cmuclmtk/0.7/cmuclmtk-0.7.src.tar.gz"
	cmuclmtk_cmd_src := exec.Command("curl", "-L", cmuclmtk_src_url, "-o", "source.tar.gz")
	err = cmuclmtk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmuclmtk_cmd_direct := exec.Command("./binary")
	err = cmuclmtk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
