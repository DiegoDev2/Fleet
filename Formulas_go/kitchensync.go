package main

// KitchenSync - Fast efficiently sync database without dumping & reloading
// Homepage: https://github.com/willbryant/kitchen_sync

import (
	"fmt"
	
	"os/exec"
)

func installKitchenSync() {
	// Método 1: Descargar y extraer .tar.gz
	kitchensync_tar_url := "https://github.com/willbryant/kitchen_sync/archive/refs/tags/v2.20.tar.gz"
	kitchensync_cmd_tar := exec.Command("curl", "-L", kitchensync_tar_url, "-o", "package.tar.gz")
	err := kitchensync_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kitchensync_zip_url := "https://github.com/willbryant/kitchen_sync/archive/refs/tags/v2.20.zip"
	kitchensync_cmd_zip := exec.Command("curl", "-L", kitchensync_zip_url, "-o", "package.zip")
	err = kitchensync_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kitchensync_bin_url := "https://github.com/willbryant/kitchen_sync/archive/refs/tags/v2.20.bin"
	kitchensync_cmd_bin := exec.Command("curl", "-L", kitchensync_bin_url, "-o", "binary.bin")
	err = kitchensync_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kitchensync_src_url := "https://github.com/willbryant/kitchen_sync/archive/refs/tags/v2.20.src.tar.gz"
	kitchensync_cmd_src := exec.Command("curl", "-L", kitchensync_src_url, "-o", "source.tar.gz")
	err = kitchensync_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kitchensync_cmd_direct := exec.Command("./binary")
	err = kitchensync_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libpq")
exec.Command("latte", "install", "libpq")
	fmt.Println("Instalando dependencia: mysql-client")
exec.Command("latte", "install", "mysql-client")
}
