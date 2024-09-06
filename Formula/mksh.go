package main

// Mksh - MirBSD Korn Shell
// Homepage: http://www.mirbsd.org/mksh.htm

import (
	"fmt"
	
	"os/exec"
)

func installMksh() {
	// Método 1: Descargar y extraer .tar.gz
	mksh_tar_url := "http://www.mirbsd.org/MirOS/dist/mir/mksh/mksh-R59c.tgz"
	mksh_cmd_tar := exec.Command("curl", "-L", mksh_tar_url, "-o", "package.tar.gz")
	err := mksh_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mksh_zip_url := "http://www.mirbsd.org/MirOS/dist/mir/mksh/mksh-R59c.tgz"
	mksh_cmd_zip := exec.Command("curl", "-L", mksh_zip_url, "-o", "package.zip")
	err = mksh_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mksh_bin_url := "http://www.mirbsd.org/MirOS/dist/mir/mksh/mksh-R59c.tgz"
	mksh_cmd_bin := exec.Command("curl", "-L", mksh_bin_url, "-o", "binary.bin")
	err = mksh_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mksh_src_url := "http://www.mirbsd.org/MirOS/dist/mir/mksh/mksh-R59c.tgz"
	mksh_cmd_src := exec.Command("curl", "-L", mksh_src_url, "-o", "source.tar.gz")
	err = mksh_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mksh_cmd_direct := exec.Command("./binary")
	err = mksh_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
