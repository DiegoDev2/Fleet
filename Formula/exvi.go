package main

// ExVi - UTF8-friendly version of traditional vi
// Homepage: https://ex-vi.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installExVi() {
	// Método 1: Descargar y extraer .tar.gz
	exvi_tar_url := "https://downloads.sourceforge.net/project/ex-vi/ex-vi/050325/ex-050325.tar.bz2"
	exvi_cmd_tar := exec.Command("curl", "-L", exvi_tar_url, "-o", "package.tar.gz")
	err := exvi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	exvi_zip_url := "https://downloads.sourceforge.net/project/ex-vi/ex-vi/050325/ex-050325.tar.bz2"
	exvi_cmd_zip := exec.Command("curl", "-L", exvi_zip_url, "-o", "package.zip")
	err = exvi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	exvi_bin_url := "https://downloads.sourceforge.net/project/ex-vi/ex-vi/050325/ex-050325.tar.bz2"
	exvi_cmd_bin := exec.Command("curl", "-L", exvi_bin_url, "-o", "binary.bin")
	err = exvi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	exvi_src_url := "https://downloads.sourceforge.net/project/ex-vi/ex-vi/050325/ex-050325.tar.bz2"
	exvi_cmd_src := exec.Command("curl", "-L", exvi_src_url, "-o", "source.tar.gz")
	err = exvi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	exvi_cmd_direct := exec.Command("./binary")
	err = exvi_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
