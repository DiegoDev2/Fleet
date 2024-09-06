package main

// Pgdbf - Converter of XBase/FoxPro tables to PostgreSQL
// Homepage: https://github.com/kstrauser/pgdbf

import (
	"fmt"
	
	"os/exec"
)

func installPgdbf() {
	// Método 1: Descargar y extraer .tar.gz
	pgdbf_tar_url := "https://downloads.sourceforge.net/project/pgdbf/pgdbf/0.6.2/pgdbf-0.6.2.tar.xz"
	pgdbf_cmd_tar := exec.Command("curl", "-L", pgdbf_tar_url, "-o", "package.tar.gz")
	err := pgdbf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgdbf_zip_url := "https://downloads.sourceforge.net/project/pgdbf/pgdbf/0.6.2/pgdbf-0.6.2.tar.xz"
	pgdbf_cmd_zip := exec.Command("curl", "-L", pgdbf_zip_url, "-o", "package.zip")
	err = pgdbf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgdbf_bin_url := "https://downloads.sourceforge.net/project/pgdbf/pgdbf/0.6.2/pgdbf-0.6.2.tar.xz"
	pgdbf_cmd_bin := exec.Command("curl", "-L", pgdbf_bin_url, "-o", "binary.bin")
	err = pgdbf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgdbf_src_url := "https://downloads.sourceforge.net/project/pgdbf/pgdbf/0.6.2/pgdbf-0.6.2.tar.xz"
	pgdbf_cmd_src := exec.Command("curl", "-L", pgdbf_src_url, "-o", "source.tar.gz")
	err = pgdbf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgdbf_cmd_direct := exec.Command("./binary")
	err = pgdbf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
