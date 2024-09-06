package main

// Pgtune - Tuning wizard for postgresql.conf
// Homepage: https://github.com/gregs1104/pgtune

import (
	"fmt"
	
	"os/exec"
)

func installPgtune() {
	// Método 1: Descargar y extraer .tar.gz
	pgtune_tar_url := "https://ftp.postgresql.org/pub/projects/pgFoundry/pgtune/pgtune/0.9.3/pgtune-0.9.3.tar.gz"
	pgtune_cmd_tar := exec.Command("curl", "-L", pgtune_tar_url, "-o", "package.tar.gz")
	err := pgtune_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pgtune_zip_url := "https://ftp.postgresql.org/pub/projects/pgFoundry/pgtune/pgtune/0.9.3/pgtune-0.9.3.zip"
	pgtune_cmd_zip := exec.Command("curl", "-L", pgtune_zip_url, "-o", "package.zip")
	err = pgtune_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pgtune_bin_url := "https://ftp.postgresql.org/pub/projects/pgFoundry/pgtune/pgtune/0.9.3/pgtune-0.9.3.bin"
	pgtune_cmd_bin := exec.Command("curl", "-L", pgtune_bin_url, "-o", "binary.bin")
	err = pgtune_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pgtune_src_url := "https://ftp.postgresql.org/pub/projects/pgFoundry/pgtune/pgtune/0.9.3/pgtune-0.9.3.src.tar.gz"
	pgtune_cmd_src := exec.Command("curl", "-L", pgtune_src_url, "-o", "source.tar.gz")
	err = pgtune_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pgtune_cmd_direct := exec.Command("./binary")
	err = pgtune_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
