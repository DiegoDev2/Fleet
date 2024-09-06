package main

// Rc - Implementation of the AT&T Plan 9 shell
// Homepage: https://doc.cat-v.org/plan_9/4th_edition/papers/rc

import (
	"fmt"
	
	"os/exec"
)

func installRc() {
	// Método 1: Descargar y extraer .tar.gz
	rc_tar_url := "https://web.archive.org/web/20200227085442/static.tobold.org/rc/rc-1.7.4.tar.gz"
	rc_cmd_tar := exec.Command("curl", "-L", rc_tar_url, "-o", "package.tar.gz")
	err := rc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rc_zip_url := "https://web.archive.org/web/20200227085442/static.tobold.org/rc/rc-1.7.4.zip"
	rc_cmd_zip := exec.Command("curl", "-L", rc_zip_url, "-o", "package.zip")
	err = rc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rc_bin_url := "https://web.archive.org/web/20200227085442/static.tobold.org/rc/rc-1.7.4.bin"
	rc_cmd_bin := exec.Command("curl", "-L", rc_bin_url, "-o", "binary.bin")
	err = rc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rc_src_url := "https://web.archive.org/web/20200227085442/static.tobold.org/rc/rc-1.7.4.src.tar.gz"
	rc_cmd_src := exec.Command("curl", "-L", rc_src_url, "-o", "source.tar.gz")
	err = rc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rc_cmd_direct := exec.Command("./binary")
	err = rc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
