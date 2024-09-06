package main

// Zsync - File transfer program
// Homepage: http://zsync.moria.org.uk/

import (
	"fmt"
	
	"os/exec"
)

func installZsync() {
	// Método 1: Descargar y extraer .tar.gz
	zsync_tar_url := "http://zsync.moria.org.uk/download/zsync-0.6.2.tar.bz2"
	zsync_cmd_tar := exec.Command("curl", "-L", zsync_tar_url, "-o", "package.tar.gz")
	err := zsync_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	zsync_zip_url := "http://zsync.moria.org.uk/download/zsync-0.6.2.tar.bz2"
	zsync_cmd_zip := exec.Command("curl", "-L", zsync_zip_url, "-o", "package.zip")
	err = zsync_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	zsync_bin_url := "http://zsync.moria.org.uk/download/zsync-0.6.2.tar.bz2"
	zsync_cmd_bin := exec.Command("curl", "-L", zsync_bin_url, "-o", "binary.bin")
	err = zsync_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	zsync_src_url := "http://zsync.moria.org.uk/download/zsync-0.6.2.tar.bz2"
	zsync_cmd_src := exec.Command("curl", "-L", zsync_src_url, "-o", "source.tar.gz")
	err = zsync_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	zsync_cmd_direct := exec.Command("./binary")
	err = zsync_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
