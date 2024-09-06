package main

// Minetest - Free, open source voxel game engine and game
// Homepage: https://www.minetest.net/

import (
	"fmt"
	
	"os/exec"
)

func installMinetest() {
	// Método 1: Descargar y extraer .tar.gz
	minetest_tar_url := "https://github.com/minetest/minetest/archive/refs/tags/5.9.0.tar.gz"
	minetest_cmd_tar := exec.Command("curl", "-L", minetest_tar_url, "-o", "package.tar.gz")
	err := minetest_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	minetest_zip_url := "https://github.com/minetest/minetest/archive/refs/tags/5.9.0.zip"
	minetest_cmd_zip := exec.Command("curl", "-L", minetest_zip_url, "-o", "package.zip")
	err = minetest_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	minetest_bin_url := "https://github.com/minetest/minetest/archive/refs/tags/5.9.0.bin"
	minetest_cmd_bin := exec.Command("curl", "-L", minetest_bin_url, "-o", "binary.bin")
	err = minetest_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	minetest_src_url := "https://github.com/minetest/minetest/archive/refs/tags/5.9.0.src.tar.gz"
	minetest_cmd_src := exec.Command("curl", "-L", minetest_src_url, "-o", "source.tar.gz")
	err = minetest_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	minetest_cmd_direct := exec.Command("./binary")
	err = minetest_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: jsoncpp")
	exec.Command("latte", "install", "jsoncpp").Run()
	fmt.Println("Instalando dependencia: libogg")
	exec.Command("latte", "install", "libogg").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
	fmt.Println("Instalando dependencia: luajit")
	exec.Command("latte", "install", "luajit").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxi")
	exec.Command("latte", "install", "libxi").Run()
	fmt.Println("Instalando dependencia: libxxf86vm")
	exec.Command("latte", "install", "libxxf86vm").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
	fmt.Println("Instalando dependencia: openal-soft")
	exec.Command("latte", "install", "openal-soft").Run()
	fmt.Println("Instalando dependencia: xinput")
	exec.Command("latte", "install", "xinput").Run()
}
