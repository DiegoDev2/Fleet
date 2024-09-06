package main

// Mmtabbarview - Modernized and view-based rewrite of PSMTabBarControl
// Homepage: https://mimo42.github.io/MMTabBarView/

import (
	"fmt"
	
	"os/exec"
)

func installMmtabbarview() {
	// Método 1: Descargar y extraer .tar.gz
	mmtabbarview_tar_url := "https://github.com/MiMo42/MMTabBarView/archive/refs/tags/v1.4.2.tar.gz"
	mmtabbarview_cmd_tar := exec.Command("curl", "-L", mmtabbarview_tar_url, "-o", "package.tar.gz")
	err := mmtabbarview_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mmtabbarview_zip_url := "https://github.com/MiMo42/MMTabBarView/archive/refs/tags/v1.4.2.zip"
	mmtabbarview_cmd_zip := exec.Command("curl", "-L", mmtabbarview_zip_url, "-o", "package.zip")
	err = mmtabbarview_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mmtabbarview_bin_url := "https://github.com/MiMo42/MMTabBarView/archive/refs/tags/v1.4.2.bin"
	mmtabbarview_cmd_bin := exec.Command("curl", "-L", mmtabbarview_bin_url, "-o", "binary.bin")
	err = mmtabbarview_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mmtabbarview_src_url := "https://github.com/MiMo42/MMTabBarView/archive/refs/tags/v1.4.2.src.tar.gz"
	mmtabbarview_cmd_src := exec.Command("curl", "-L", mmtabbarview_src_url, "-o", "source.tar.gz")
	err = mmtabbarview_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mmtabbarview_cmd_direct := exec.Command("./binary")
	err = mmtabbarview_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
