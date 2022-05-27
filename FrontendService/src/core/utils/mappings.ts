export const createEntitiesMap = (entities: any, keySelector = (entity: any) => entity.id, valueSelector = (entity: any) => entity) => {
	return entities?.reduce((map: any, entity: any) => {
		map[keySelector(entity)] = valueSelector(entity);
		return map;
	}, {}) ?? {};
}