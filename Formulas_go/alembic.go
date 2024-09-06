package main

// Alembic - Open computer graphics interchange framework
// Homepage: http://alembic.io

import (
	"fmt"
	
	"os/exec"
)

func installAlembic() {
	// Método 1: Descargar y extraer .tar.gz
	alembic_tar_url := "https://github.com/alembic/alembic/archive/refs/tags/1.8.6.tar.gz"
	alembic_cmd_tar := exec.Command("curl", "-L", alembic_tar_url, "-o", "package.tar.gz")
	err := alembic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	alembic_zip_url := "https://github.com/alembic/alembic/archive/refs/tags/1.8.6.zip"
	alembic_cmd_zip := exec.Command("curl", "-L", alembic_zip_url, "-o", "package.zip")
	err = alembic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	alembic_bin_url := "https://github.com/alembic/alembic/archive/refs/tags/1.8.6.bin"
	alembic_cmd_bin := exec.Command("curl", "-L", alembic_bin_url, "-o", "binary.bin")
	err = alembic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	alembic_src_url := "https://github.com/alembic/alembic/archive/refs/tags/1.8.6.src.tar.gz"
	alembic_cmd_src := exec.Command("curl", "-L", alembic_src_url, "-o", "source.tar.gz")
	err = alembic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	alembic_cmd_direct := exec.Command("./binary")
	err = alembic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: hdf5")
exec.Command("latte", "install", "hdf5")
	fmt.Println("Instalando dependencia: imath")
exec.Command("latte", "install", "imath")
	fmt.Println("Instalando dependencia: libaec")
exec.Command("latte", "install", "libaec")
}
